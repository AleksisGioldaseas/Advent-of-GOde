package main

import "testing"

func BenchmarkMyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Call the function you want to benchmark here
		slow()
	}
}
