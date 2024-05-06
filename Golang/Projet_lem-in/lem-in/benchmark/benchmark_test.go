package main

import (
	"testing"

	internal "lem-in/internal"
)

// command : go test -bench . -benchmem

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var args []string
		args = append(args, "../example/example07.txt")
		internal.Travail(args)
	}
}
