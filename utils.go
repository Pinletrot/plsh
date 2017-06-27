package plsh

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
)

type hashParam struct {
	a, b, p uint32
}

type hashParams []hashParam

func (hp *hashParam) hash(x uint32) uint32 {
	return (hp.a*x + hp.b) % hp.p
}

func MD5HashByte(x []byte) (ret uint32) {
	f := md5.New()
	f.Write(x)
	b := f.Sum(nil)
	b_buf := bytes.NewBuffer(b)

	binary.Read(b_buf, binary.BigEndian, &ret)

	return
}

func MD5HashString(x string) (ret uint32) {
	f := md5.New()
	f.Write([]byte(x))
	b := f.Sum(nil)
	b_buf := bytes.NewBuffer(b)

	binary.Read(b_buf, binary.BigEndian, &ret)

	return
}

func Uint32ToBytes(i uint32) (ret []byte) {
	ret = make([]byte, 4)
	binary.BigEndian.PutUint32(ret, i)
	return
}

func BytesToUint32(buf []byte) (ret uint32) {
	return uint32(binary.BigEndian.Uint32(buf))
}
