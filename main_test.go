package main

import (
	"testing"
)

func BenchmarkGetHashUsingArray10Million(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		firstValue := "abc"
		hs := HS{"", make(HashArray, 10_000_001)}
		hs.HashList[0] = firstValue
		hs.GetHashUsingArray()
		if hs.LastHash != "bf34d93b4be2a313b06cdf9d805c5f3d140abd872c37199701fb1e43fe479923" {
			b.Error("Unexpected Result: " + hs.LastHash)
		}
	}
}

func BenchmarkGetHashUsingLoop10Million(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		firstValue := "abc"
		result := GetHashUsingLoop(firstValue, 10_000_000)
		if result != "bf34d93b4be2a313b06cdf9d805c5f3d140abd872c37199701fb1e43fe479923" {
			b.Error("Unexpected result: " + result)
		}
	}
}
