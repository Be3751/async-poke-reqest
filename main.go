package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/be3/async-poke-reqest/conf"
	myhttp "github.com/be3/async-poke-reqest/http"
	"github.com/be3/async-poke-reqest/model"
	"github.com/be3/async-poke-reqest/util"
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

func main() {
	flag.Parse()
	flags := flag.Args()

	// 並列処理か逐次処理のどちらかを指定するオプション
	option := flags[0]
	// 実行回数
	N, err := strconv.Atoi(flags[1])
	if err != nil {
		AtoiError := fmt.Errorf("tried to convert string %s into int value but an error occured: %w", flags[1], err)
		fmt.Println(AtoiError)
		os.Exit(1)
	}

	var avgRuntime time.Duration

	switch option {
	case "p":
		avgRuntime = util.CalcAvgRuntime(paraGetAllPokemon, N)
	case "s":
		avgRuntime = util.CalcAvgRuntime(getAllPokemon, N)
	default:
		fmt.Println("Invalid flag! Use 'p' or 's' as a flag.")
		return
	}

	fmt.Printf("It took %s on average.\n", avgRuntime)
}
