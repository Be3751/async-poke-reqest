package main

import (
	"testing"
)

func BenchmarkGetAllPokemon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAllPokemon()
	}
}

func BenchmarkAsyncGetAllPokemon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		asyncGetAllPokemon()
	}
}
