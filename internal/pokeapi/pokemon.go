package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ekaloi/pokedex/poketypes"
)

func (c *Client) GetPokemon(pokemon string) (poketypes.PokemonActual, error){
	url := baseURL + "/pokemon/" + pokemon

	data, present := c.Cache.Get(url)
	var result poketypes.PokemonActual
	if !present{

		req, err := http.NewRequest("GET", url,nil)
		if err != nil{
			return poketypes.PokemonActual{}, fmt.Errorf("error creating req: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil{
			return poketypes.PokemonActual{}, fmt.Errorf("error creating req: %w", err)
		}

		defer res.Body.Close()

		dat, err := io.ReadAll(res.Body)
		c.Cache.Add(url, dat)
		if err != nil {
			return poketypes.PokemonActual{}, fmt.Errorf("error creating req: %w", err)
		}
		err = json.Unmarshal(dat, &result)
		if err != nil{
			return poketypes.PokemonActual{}, fmt.Errorf("error creating req: %w", err)
		}
	}else{
		err := json.Unmarshal(data, &result)
		if err != nil {
			return  poketypes.PokemonActual{}, fmt.Errorf("error creating req: %w", err)
		}
	}

	return result, nil
}
	

	


	
