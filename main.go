package main

import (
	"crypto/sha256"
	"fmt"
)

type HASHSTRUCT struct {
	s string
	// buf Buffer
	b []byte
	d [32]byte
}

const hextable = "0123456789abcdef"

func (HS *HASHSTRUCT) EncodeToHexBytes() {
	HS.b = make([]byte, 64)
	j := 0
	for _, v := range HS.d {
		HS.b[j] = hextable[v>>4]
		HS.b[j+1] = hextable[v&0x0f]
		j += 2
	}
}

func (HS *HASHSTRUCT) HashFn() {
	HS.d = sha256.Sum256(HS.b)
}

func (HS *HASHSTRUCT) HashLoop(loops int) {
	HS.b = []byte(HS.s)
	for i := 0; i < loops; i++ {
		HS.HashFn()
		HS.EncodeToHexBytes()
	}
}

// func (h *HashThis) HashLoop(loops int) {
// 	var buf bytes.Buffer

// 	buf.WriteString(h.inputStr)
// 	// hash.Write(buf.Bytes())
// 	hash.Sum(buf.Bytes())
// 	// h.digest = sha256.Sum256(buf.Bytes())
// 	buf.Reset()
// 	// h.hashStr = hex.EncodeToString(h.digest[:])
// 	// for i := 0; i < loops-1; i++ {
// 	// 	h.digest = sha256.Sum256([]byte(h.hashStr))
// 	// 	h.hashStr = hex.EncodeToString(h.digest[:])
// 	// }
// }

func main() {

	loops := 2

	HS := &HASHSTRUCT{s: "abc"}
	// HS.s = "a65e2e1666763cceac4458e8bd26e0134dfa09178c8b106521265fe279c2e85c"
	HS.HashLoop(loops)

	// HashLoop(result, loops)

	fmt.Printf("%x", HS.b)

	// hash := sha256.New()

	// buf.WriteString("abc")
	// hash.Write(buf.Bytes())
	// next := hex.EncodeToString(hash.Sum(nil))
	// buf.Reset()
	// hash.Reset()
	// for i := 1; i < loops; i++ {
	// 	buf.WriteString(next)
	// 	hash.Write(buf.Bytes())
	// 	next = hex.EncodeToString(hash.Sum(nil))
	// 	buf.Reset()
	// 	hash.Reset()
	// }
	// fmt.Printf("%s\n", next)

	// argv := os.Args
	// argc := len(argv)
	// command := argv[0]
	// if argc < 2 {
	// 	fmt.Printf("Usage:\n\t command <string> [number of loops Default = 100Million]\nExample:\n\t%s abc\n\n\t%s abc 100\n", command, command)
	// 	os.Exit(1)
	// }
	// start := time.Now()

	// input := argv[1]
	// loops := 100_000_000
	// if argc > 2 {
	// 	num, err := strconv.Atoi(argv[2])
	// 	if err != nil {
	// 		fmt.Printf("Received 2nd arg but couldnt convert it to int!")
	// 		os.Exit(1)
	// 	}
	// 	loops = num
	// }

	// hash := HashLoop(input, loops)

	// elapsed := time.Since(start)
	// fmt.Printf("%s,%s\n\ncompleted in: %s\n", input, hash, elapsed)
}
