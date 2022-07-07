package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/be3/async-poke-reqest/model"
)

func DecodeToPokemon(resp *http.Response) (model.Pokemon, error) {
	var pokemon model.Pokemon

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return pokemon, err
	}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		fmt.Println(err)
		return pokemon, err
	}

	return pokemon, nil
}
