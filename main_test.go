package main

import (
	"testing"
)

// func TestHasherWithEmptyString(t *testing.T) {
// 	result := HashLoop("", 1)
// 	if result != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
// 		t.Fatalf("TestHasherWithEmptyString: FAILED. Unexpected result: %s", result)
// 	}
// }

// func TestHasherWith0Loops(t *testing.T) {
// 	result := HashLoop("abc", 0)
// 	if result != "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad" {
// 		t.Fatalf("TestHasherWithEmptyString: FAILED. Unexpected result: %s", result)
// 	}
// }

// func BenchmarkGetHashUsingLoop10Million(b *testing.B) {
// 	b.ReportAllocs()

// 	HS := &HASHSTRUCT{s: "abc"}

// 	for i := 0; i < b.N; i++ {
// 		HS.HashLoop(10_000_000)
// 		result := string(HS.b)
// 		if result != "bf34d93b4be2a313b06cdf9d805c5f3d140abd872c37199701fb1e43fe479923" {
// 			b.Error("Unexpected result: " + result)
// 		}
// 	}
// }

func BenchmarkGetHashUsingLoop10Million(b *testing.B) {
	b.ReportAllocs()
	h := &Hasher{bs: []byte("abc"), b: make([]byte, 64)}
	for i := 0; i < b.N; i++ {
		h.HashFn()
		h.EncodeToHexBytes()
	}
}

// func BenchmarkGetHashUsingLoop10Million(b *testing.B) {
// 	b.ReportAllocs()
// 	var result string
// 	var firstValue string
// 	for i := 0; i < b.N; i++ {
// 		firstValue = "abc"
// 		result = HashLoop(firstValue, 10_000_000)
// 		if result != "bf34d93b4be2a313b06cdf9d805c5f3d140abd872c37199701fb1e43fe479923" {
// 			b.Error("Unexpected result: " + result)
// 		}
// 	}
// }
