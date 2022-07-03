package main

import (
	"testing"
)

func BenchmarkGetPokemon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getPokemon()
	}
}
