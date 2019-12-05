package m

import (
	"fmt"
	"testing"
)

var p = Policy{
	AS{"13800010002", "u1", "g1"},
	AO{"D100010001", "00:11:22:33:44:55"},
	1,
	AE{1575468182, 1576468182, "*.*.*.*"},
}

func Test_ToBytes(t *testing.T) {
	fmt.Printf("%s\n", p.ToBytes())
}

func Test_GetID(t *testing.T) {
	println(p.GetID())
}
