package http

import (
	"net/http"
	"strconv"

	"github.com/be3/async-poke-reqest/conf"
	"github.com/be3/async-poke-reqest/model"
	"github.com/be3/async-poke-reqest/util"
)

func GetPokemon(number int) model.Pokemon {
	url := conf.POKEMON_API + strconv.Itoa(number)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	pokemon, err := util.DecodeToPokemon(resp)
	return pokemon
}
