package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMinHash(t *testing.T) {
	rand.Seed(10)

	var n uint32 = 3
	var i uint32

	mhp := make(hashParams, n)
	for i = 0; i < 3; i++ {
		mhp[i] = hashParam{rand.Uint32(), 1, (1 << 31) - 1}
	}
	mhv := NewMinHashValue(n, mhp)

	for i = 0; i < 3; i++ {
		t.Log(mhv.values)
		mhv.Update(fmt.Sprint(i))
	}

}
