package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	Nmae string `json:"name"`
}

func getPokemon() Post {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/10")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var post Post

	if err := json.Unmarshal(body, &post); err != nil {
		panic(err)
	}

	return post
}

func main() {
	// fmt.Println(resp.Status)
	// fmt.Println(string(body))
	post := getPokemon()
	fmt.Println(post)
}
