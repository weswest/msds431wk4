package main

import "testing"

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// It's okay to ignore the error in a benchmark
		_ = run("housesInput.csv", "housesOutputGo.txt", 100)
	}
}
