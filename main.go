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

func (h *HS) GetHashUsingArray() {
	hashit := func(s string) string {
		digest := sha256.Sum256([]byte(s))
		return hex.EncodeToString(digest[:])
	}
	hl := h.HashList
	for i := 1; i < len(hl); i++ {
		(hl)[i] = hashit((hl)[i-1])
	}
	h.LastHash = hl[len(hl)-1]
}

func GetHashUsingLoop(s string, loops int) string {
	hashit := func(s *string) {
		digest := sha256.Sum256([]byte(*s))
		*s = hex.EncodeToString(digest[:])
	}
	hash := s
	for i := 0; i < loops; i++ {
		hashit(&hash)
	}
	return hash
}

func main() {}
