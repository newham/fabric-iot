package pow

import (
	"encoding/json"

	"github.com/newham/goblockchain/core"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/newham/fabric-iot/test/consensus/kafka"
)

const (
	TOPIC = "/test/pow"
	HOST  = "tcp://localhost:1883"
)

func PoW(n, nBit int) {
	ss := make(chan bool, n)
	sg := make(chan bool)
	done := make(chan string, 1)
	for i := 0; i < n; i++ {
		go NewNode(TOPIC, HOST, n, nBit).Mine(ss, sg, done)
	}
	//now wait all node started
	for i := 0; i < n; i++ {
		<-ss
	}
	// log.Println("All node started")
	//let all node to submit Transaction
	close(sg)
	//wait finish
	// log.Println("Winner is:", <-done)
}

type Msg struct {
	ID    string
	Block core.Block
}

func (m Msg) Bytes() []byte {
	bts, err := json.Marshal(m)
	if err != nil {
		println(err.Error())
	}
	return bts
}

type Node struct {
	topic string
	mc    *kafka.MQTTCli
	cNum  int
	cfNum int
	nBit  int
}

func NewNode(topic string, host string, cNum, nBit int) *Node {
	return &Node{topic, kafka.NewMQTTCli(host), cNum, 0, nBit}
}

func (n *Node) Mine(ss, sg chan bool, done chan string) {
	n.mc.Sub(n.topic+"/mine", func(client mqtt.Client, message mqtt.Message) {
		//receive other's block
		// core.NewProofOfWork(core.NewBlock(nil, nil, nBit))
		b := message.Payload()
		msg := &Msg{}
		err := json.Unmarshal(b, msg)
		if err != nil {
			// log.Println(err.Error())
		}
		ok, _ := core.NewProofOfWork(&msg.Block).Validate(msg.Block.Nonce)
		if ok {
			n.mc.Pub(n.topic+"/to/"+msg.ID, "ok")
		}
	})
	n.mc.Sub(n.topic+"/to/"+n.mc.ID(), func(client mqtt.Client, message mqtt.Message) {
		if string(message.Payload()) == "ok" {
			n.cfNum++
		}
		if n.cfNum > n.cNum/2 {
			done <- n.mc.ID()
		}
	})
	ss <- true
	// log.Println("Node", n.mc.ID(), "started")
	if _, ok := <-sg; !ok {
		nb := core.NewGenesisBlock(n.mc.ID(), n.nBit)
		//start to pow
		nonce, _ := core.NewProofOfWork(nb).Work()
		nb.Nonce = nonce
		// log.Println("Node", n.mc.ID(), "minded:", nonce)
		msg := Msg{n.mc.ID(), *nb}
		//pub to all
		n.mc.Pub(n.topic+"/mine", msg.Bytes())
	}
}
