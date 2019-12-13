package kafka

import (
	"fmt"
	"testing"
	"time"
)

func TestKafka(t *testing.T) {
	for j := 10; j <= 500; j += 20 {
		start := time.Now().UnixNano()
		Kafka(j)
		end := time.Now().UnixNano()
		fmt.Printf("%-4d%.2f\n", j, float32(end-start)/1e6)
	}

}
