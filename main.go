package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	POKEMON_API = "https://pokeapi.co/api/v2/pokemon/"
	POKEMON_SUM = 151
	COUNT       = 1
)

type Pokemon struct {
	Name string `json:"name"`
}

func getPokemon(number int) Pokemon {
	url := POKEMON_API + strconv.Itoa(number)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var pokemon Pokemon

	if err := json.Unmarshal(body, &pokemon); err != nil {
		panic(err)
	}

	return pokemon
}

func getAllPokemon() []Pokemon {
	var pokemons []Pokemon

	for i := 0; i < POKEMON_SUM; i++ {
		number := i + 1
		pokemon := getPokemon(number)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons
}

func asyncRequest(url string, c chan []byte) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	c <- body
}

func asyncGetPokemon(number int, c chan Pokemon, wg *sync.WaitGroup) {
	defer wg.Done()

	url := POKEMON_API + strconv.Itoa(number)
	ch := make(chan []byte)
	go asyncRequest(url, ch)
	body := <-ch

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		panic(err)
	}

	// APIから取得したデータをチャネルに格納
	c <- pokemon
}

func asyncGetAllPokemon() []Pokemon {
	var pokemons []Pokemon
	var wg sync.WaitGroup

	for i := 0; i < POKEMON_SUM; i++ {
		// Pokemon型のチャネルを生成
		c := make(chan Pokemon)

		// WaitGroupにGoルーチンを追加
		wg.Add(1)
		number := i + 1
		go asyncGetPokemon(number, c, &wg)
		pokemon := <-c
		pokemons = append(pokemons, pokemon)
	}
	wg.Wait()

	return pokemons
}

func main() {
	start := time.Now()
	var pokemons []Pokemon
	for i := 0; i < COUNT; i++ {
		pokemons = getAllPokemon()
	}
	fmt.Println(pokemons)
	end := time.Now()
	fmt.Printf("It took %s.\n", end.Sub(start))

	start = time.Now()
	for i := 0; i < COUNT; i++ {
		pokemons = asyncGetAllPokemon()
	}
	fmt.Println(pokemons)
	end = time.Now()
	fmt.Printf("It took %s.\n", end.Sub(start))
}
