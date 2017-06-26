package main

import (
	"math/rand"
)

type LSH struct {
	seed          int64
	p             uint32
	nBands        uint32
	nRowsPerBand  uint32
	nMinHashFunc  uint32
	minHashParams hashParams
	hashParams    hashParams
}

func NewLSH(seed int64, p, nBands, nRowsPerBand uint32) (lsh *LSH) {
	lsh = &LSH{}
	lsh.seed = seed
	lsh.p = p
	lsh.nBands = nBands
	lsh.nRowsPerBand = nRowsPerBand
	lsh.nMinHashFunc = nBands * nRowsPerBand

	lsh.minHashParams = make(hashParams, lsh.nMinHashFunc)
	lsh.hashParams = make(hashParams, lsh.nBands)

	var i uint32
	rand.Seed(lsh.seed)
	for i = 0; i < lsh.nBands; i++ {
		lsh.hashParams[i] = hashParam{rand.Uint32(), rand.Uint32(), p}
	}
	for i = 0; i < lsh.nMinHashFunc; i++ {
		lsh.minHashParams[i] = hashParam{rand.Uint32(), rand.Uint32(), p}
	}

	return
}

func (lsh *LSH) GetBucketIds(msv *MinHashValue) []uint32 {
	var i, j, k uint32
	var buf = make([]byte, 4*lsh.nRowsPerBand)

	bucketIds := make([]uint32, lsh.nBands)
	for i = 0; i < lsh.nBands; i++ {
		for j = 0; j < lsh.nRowsPerBand; j++ {
			v := Uint32ToBytes(msv.values[i*lsh.nRowsPerBand+j])
			for k = 0; k < 4; k++ {
				buf[4*j+k] = v[k]
			}
		}
		t := MD5HashByte(buf)
		bucketIds[i] = lsh.hashParams[i].hash(t)
	}
	return bucketIds
}
