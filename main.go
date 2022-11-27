package main

import (
	// "crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/minio/sha256-simd"
)

type HASHSTRUCT struct {
	s string
	// buf bytes.Buffer
	bs []byte
	b  []byte
	d  [32]byte
}

const hextable = "0123456789abcdef"

func (HS *HASHSTRUCT) EncodeToHexBytes() {
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
	for i := 0; i < loops; i++ {
		HS.HashFn()
		HS.EncodeToHexBytes()
	}
}

func main() {
	argv := os.Args
	argc := len(argv)
	command := argv[0]
	if argc < 2 {
		fmt.Printf("Usage:\n\t command <string> [number of loops Default = 100Million]\nExample:\n\t%s abc\n\n\t%s abc 100\n", command, command)
		os.Exit(1)
	}
	start := time.Now()

	input := argv[1]
	loops := 100_000_000
	if argc > 2 {
		num, err := strconv.Atoi(argv[2])
		if err != nil {
			fmt.Printf("Received 2nd arg but couldnt convert it to int!")
			os.Exit(1)
		}
		loops = num
	}

	HS := &HASHSTRUCT{s: input}
	HS.bs = []byte(HS.s)
	HS.b = make([]byte, 64)
	HS.HashLoop(loops)

	elapsed := time.Since(start)
	fmt.Printf("%s,%s\n\ncompleted in: %s\n", input, HS.b, elapsed)

}
