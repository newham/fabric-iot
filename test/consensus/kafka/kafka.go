package kafka

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	TOPIC = "/test/kafka"
	HOST  = "tcp://localhost:1883"
)

func Kafka(n int) {

	sn := make(chan bool)
	//start order
	go NewOrder(TOPIC, HOST, n).Rank(sn)
	<-sn
	//log.Println("Order started")
	//start peer
	ss := make(chan bool, n)
	sg := make(chan bool)
	for i := 0; i < n; i++ {
		go NewPeer(TOPIC, HOST).Submit(ss, sg)
	}
	//now wait all peer started
	for i := 0; i < n; i++ {
		<-ss
	}
	//log.Println("All peer started")
	//let all peer to submit Transaction
	close(sg)
	//wait order finish
	<-sn
}

func Wait() {
	sn := make(chan int)
	<-sn
}

/************* MQTTCli *************/
type MQTTCli struct {
	id string
	c  mqtt.Client
}

func NewMQTTCli(host string) *MQTTCli {
	mc := &MQTTCli{}
	mc.Connect(host)
	return mc
}

func (mc *MQTTCli) ID() string {
	return mc.id
}
func (mc *MQTTCli) Pub(topic string, msg interface{}) {
	// //log.Println("pub:", topic, msg)
	token := mc.c.Publish(topic, 1, false, msg)
	if token.Error() != nil {
		//log.Println(token.Error().Error())
	}
}

func (mc *MQTTCli) Sub(topic string, f func(client mqtt.Client, message mqtt.Message)) {
	if token := mc.c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func (mc *MQTTCli) Connect(host string) {
	mc.id = strconv.FormatInt(time.Now().UnixNano(), 10)
	opts := mqtt.NewClientOptions().AddBroker(host).SetClientID(mc.id)
	mc.c = mqtt.NewClient(opts)
	if token := mc.c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

/************* Peer *************/
type Peer struct {
	topic string
	mc    *MQTTCli
}

func (p *Peer) Submit(ss, sg chan bool) {
	p.mc.Sub(p.topic+"/call", func(client mqtt.Client, message mqtt.Message) {
		//sleep cost time
		CostTime()
		if string(message.Payload()) == "confirm" {
			p.mc.Pub(p.topic+"/back", "ok")
		}
	})
	ss <- true
	//log.Println("Peer", p.mc.id, "started")
	// all peer have started, now can pub
	if _, ok := <-sg; !ok {
		p.mc.Pub(p.topic, NewTransaction(p.mc.id).Bytes())
	}

}

func NewPeer(topic string, host string) *Peer {
	return &Peer{topic, NewMQTTCli(host)}
}

/************* Order *************/
type Order struct {
	topic string
	mc    *MQTTCli
	cNum  int
	lock  bool
}

func NewOrder(topic string, host string, cNum int) *Order {
	return &Order{topic: topic, mc: NewMQTTCli(host), cNum: cNum, lock: false}
}

func (o *Order) Rank(c chan bool) {
	o.mc.Sub(o.topic, func(client mqtt.Client, message mqtt.Message) {
		GetTransaction(message.Payload())
		if o.lock {
			//log.Println("Order rank:", t.PID, "at", t.Timestamp, "->reject")
			return
		}
		//only recesive 1 submit
		o.lock = true
		//log.Println("Order rank:", t.PID, "at", t.Timestamp, "->receive")
		//sent to other clients
		o.mc.Pub(o.topic+"/call", "confirm")
	})
	//log.Println("Order sub:", o.topic)
	o.mc.Sub(o.topic+"/back", func(client mqtt.Client, message mqtt.Message) {
		if b := message.Payload(); b != nil && string(b) == "ok" {
			CostTime()
			//log.Println("Order received:", string(b))
			o.cNum -= 1
		}
		if o.cNum == 0 {
			c <- true
		}
	})
	//log.Println("Order sub:", o.topic+"/back")
	c <- true
}

/************* Transaction *************/
type Transaction struct {
	PID       string
	Timestamp int64
}

func GetTransaction(b []byte) *Transaction {
	t := &Transaction{}
	if err := json.Unmarshal(b, t); err != nil {
		return nil
	}
	return t
}

func (t *Transaction) Bytes() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		return nil
	}
	return b
}

func NewTransaction(pid string) *Transaction {
	return &Transaction{pid, time.Now().Unix()}
}

func CostTime() {
	time.Sleep(time.Microsecond * 500)
}
