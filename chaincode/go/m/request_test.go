package m

import (
	"fmt"
	"testing"
)

func Test_ABACRequest(t *testing.T) {
	r := ABACRequest{
		AS{"13800010002", "u1", "g1"},
		AO{"D100010001", "00:11:22:33:44:55"},
	}
	fmt.Printf("%s\n", r.ToBytes())
}
