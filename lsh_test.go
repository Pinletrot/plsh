package main

import (
	"testing"
)

func TestLSH_FindBucket(t *testing.T) {
	var seed int64
	var p, nBands, nRowsPerBand uint32

	seed = 10
	p = 65537
	nBands = 2
	nRowsPerBand = 4
	lsh := NewLSH(seed, p, nBands, nRowsPerBand)

	m1 := NewMinHashValue(lsh.nMinHashFunc, lsh.minHashParams)
	m2 := NewMinHashValue(lsh.nMinHashFunc, lsh.minHashParams)
	m3 := NewMinHashValue(lsh.nMinHashFunc, lsh.minHashParams)
	m4 := NewMinHashValue(lsh.nMinHashFunc, lsh.minHashParams)

	m1.Update("test")
	m2.Update("test")

	m1.Update("sss")
	m2.Update("sss")

	m2.Update("kill")
	m3.Update("kill")
	m4.Update("kill")

	t.Log(lsh.GetBucketIds(m1))
	t.Log(lsh.GetBucketIds(m2))
	t.Log(lsh.GetBucketIds(m3))
	t.Log(lsh.GetBucketIds(m4))
}

func TestLSH_FindBucket_Similarity(t *testing.T) {
	var seed int64
	var p, nBands, nRowsPerBand uint32

	seed = 10
	p = 65537
	nBands = 16
	nRowsPerBand = 16

	nGram := 2
	lsh := NewLSH(seed, p, nBands, nRowsPerBand)

	sentences := [...]string{
		"今天去哪里玩",
		"今天去这里玩",
		"今天不高兴",
		"今天超高兴"}

	m := make([]*MinHashValue, len(sentences))
	for i := 0; i < len(sentences); i++ {
		m[i] = NewMinHashValue(lsh.nMinHashFunc, lsh.minHashParams)
		for j := 0; j < len(sentences[i]) - nGram; j++ {
			m[i].Update(sentences[i][j: j + nGram])
		}
		t.Log(sentences[i], lsh.GetBucketIds(m[i]))
	}
}