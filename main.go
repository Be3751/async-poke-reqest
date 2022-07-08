package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
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
		fmt.Print("id: " + strconv.Itoa(pokemon.Id) + ", ")
		fmt.Println("name: " + pokemon.Name)
	}
	fmt.Println("-------------------")
}

func paraGetAllPokemon() {
	fmt.Println("-------------------")

	// Pokemon型のチャネルを生成
	pokeChan := make(chan model.Pokemon, conf.POKEMON_SUM)
	var wg sync.WaitGroup

	for i := 0; i < conf.POKEMON_SUM; i++ {
		wg.Add(1)
		number := i + 1
		go myhttp.ParaGetPokemon(number, pokeChan)
		wg.Done()
	}

	// 151匹のポケモンデータを取得し終えるまで処理を待機
	wg.Wait()

	cnt := 0
	for pokemon := range pokeChan {
		cnt++
		fmt.Print("id: " + strconv.Itoa(pokemon.Id) + ", ")
		fmt.Println("name: " + pokemon.Name)
		if cnt == conf.POKEMON_SUM {
			break
		}
	}

	fmt.Println("-------------------")
}

func measurer(fnc func()) time.Duration {
	fmt.Println("Let's get started to capture all pokemons!")
	start := time.Now()
	fnc()
	end := time.Now()
	return end.Sub(start)
}

func main() {
	flag.Parse()
	flag := flag.Arg(0)

	var duration time.Duration

	switch flag {
	case "p":
		duration = measurer(paraGetAllPokemon)
	case "s":
		duration = measurer(getAllPokemon)
	default:
		fmt.Println("Invalid flag! Use 'p' or 's' as a flag.")
		return
	}

	fmt.Printf("It took %s.\n", duration)
}
