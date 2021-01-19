package pow

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/newham/goblockchain/core"
)

func TestMine(t *testing.T) {
	difficulty := 10
	start := time.Now().UnixNano()
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	nb := core.NewGenesisBlock(id, difficulty)
	//start to pow
	nonce, hash := core.NewProofOfWork(nb).Work()
	end := time.Now().UnixNano()
	fmt.Printf("difficulty=%-3d id=%s nonce=%-10d hash=%x time(s)=%.6f\n", difficulty, id, nonce, hash, float32(end-start)/1e9)
}

func TestPoW(t *testing.T) {
	nBit := 16
	for j := 5; j <= 100; j += 5 {
		start := time.Now().UnixNano()
		PoW(j, nBit)
		end := time.Now().UnixNano()
		fmt.Printf("%.2f\n", float32(end-start)/1e6)
	}
}

func TestNBit(t *testing.T) {
	for i := 1; i <= 23; i++ {
		start := time.Now().UnixNano()
		id := strconv.FormatInt(time.Now().UnixNano(), 10)
		nb := core.NewGenesisBlock(id, i)
		//start to pow
		nonce, hash := core.NewProofOfWork(nb).Work()
		end := time.Now().UnixNano()
		fmt.Printf("%-4d%s %-10d %x %.6f\n", i, id, nonce, hash, float32(end-start)/1e9)
	}

}
