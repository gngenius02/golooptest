package main

import (
	"crypto/sha256"
	"fmt"
)

type HASHSTRUCT struct {
	s string
	// buf Buffer
	bs []byte
	b  []byte
	d  [32]byte
}

const hextable = "0123456789abcdef"

func (HS *HASHSTRUCT) EncodeToHexBytes() {
	// HS.b = make([]byte, 64)
	j := 0
	for _, v := range HS.d {
		HS.b[j] = hextable[v>>4]
		HS.b[j+1] = hextable[v&0x0f]
		j += 2
	}
}

func (HS *HASHSTRUCT) HashFn() {
	if HS.bs != nil {
		HS.d = sha256.Sum256(HS.bs)
		HS.bs = nil
	} else {
		HS.d = sha256.Sum256(HS.b)
	}
}

func (HS *HASHSTRUCT) HashLoop(loops int) {
	HS.bs = []byte(HS.s)
	HS.b = make([]byte, 64)
	for i := 0; i < loops; i++ {
		HS.HashFn()
		HS.EncodeToHexBytes()
	}
}

func main() {

	loops := 2

	HS := &HASHSTRUCT{s: "abc"}
	HS.HashLoop(loops)
	fmt.Printf("%s", HS.b)
}
