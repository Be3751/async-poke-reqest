package main

import (
	"fmt"
	"time"

	"github.com/be3/async-poke-reqest/conf"
	myhttp "github.com/be3/async-poke-reqest/http"
	"github.com/be3/async-poke-reqest/model"
)

func getAllPokemon() {
	fmt.Println("-------------------")
	for i := 0; i < conf.POKEMON_SUM; i++ {
		number := i + 1
		pokemon := myhttp.GetPokemon(number)
		fmt.Println(pokemon.Name)
	}
	fmt.Println("-------------------")
}

func asyncGetAllPokemon() {
	// var wg sync.WaitGroup
	fmt.Println("-------------------")

	// Pokemon型のチャネルを生成
	c := make(chan model.Pokemon, conf.POKEMON_SUM)

	for i := 0; i < conf.POKEMON_SUM; i++ {

		// WaitGroupにGoルーチンを追加
		// wg.Add(1)
		number := i + 1
		go myhttp.AsyncGetPokemon(number, c)
		pokemon := <-c
		fmt.Println(pokemon.Name)
	}

	// wg.Wait()
	fmt.Println("-------------------")
}

func main() {
	fmt.Println("Get started to request!")
	start := time.Now()
	// getAllPokemon()
	asyncGetAllPokemon()
	end := time.Now()
	fmt.Printf("It took %s.\n", end.Sub(start))
}
