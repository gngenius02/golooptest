package main

import (
	"crypto/sha256"
	"encoding/hex"
)

type HashArray []string

type HS struct {
	LastHash string
	HashList HashArray
}

func hashit(s string) string {
	digest := sha256.Sum256([]byte(s))
	return hex.EncodeToString(digest[:])
}

func (h *HS) GetHashUsingArray() {
	hl := h.HashList
	for i := 1; i < len(hl); i++ {
		(hl)[i] = hashit((hl)[i-1])
	}
	h.LastHash = hl[len(hl)-1]
}

func GetHashUsingLoop(s string, loops int) string {
	hash := hashit(s)
	for i := 0; i < loops-1; i++ {
		hash = hashit(hash)
	}
	return hash
}

func main() {}
