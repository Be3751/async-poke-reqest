package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/be3/async-poke-reqest/conf"
	"github.com/be3/async-poke-reqest/model"
)

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

func AsyncGetPokemon(number int, c chan model.Pokemon) {
	url := conf.POKEMON_API + strconv.Itoa(number)
	ch := make(chan []byte)
	go asyncRequest(url, ch)
	body := <-ch

	var pokemon model.Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		panic(err)
	}

	// APIから取得したデータをチャネルに格納
	c <- pokemon
}
