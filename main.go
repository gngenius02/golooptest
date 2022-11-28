package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/minio/sha256-simd"
)

type Hasher struct {
	bs []byte
	b  []byte
	d  [32]byte
}

const hextable = "0123456789abcdef"

func (h *Hasher) EncodeToHexBytes() {
	for i, v := range h.d {
		h.b[2*i] = hextable[v>>4]
		h.b[2*i+1] = hextable[v&0x0f]
	}
}

func (h *Hasher) HashFn() {
	if h.bs != nil {
		h.d = sha256.Sum256(h.bs)
		h.bs = nil
	} else {
		h.d = sha256.Sum256(h.b)
	}
	h.EncodeToHexBytes()
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
			fmt.Printf("ERROR: Received 2nd arg but couldnt convert it to int!\n\n")
			fmt.Printf("Usage:\n\t command <string> [number of loops Default = 100Million]\nExample:\n\t%s abc\n\n\t%s abc 100\n", command, command)
			os.Exit(1)
		}
		loops = num
	}

	h := &Hasher{bs: []byte(input), b: make([]byte, 64)}

	for i := 0; i < loops; i++ {
		h.HashFn()
	}

	elapsed := time.Since(start)
	fmt.Printf("%s,%s\n\ncompleted in: %s\n", input, h.b, elapsed)

}
