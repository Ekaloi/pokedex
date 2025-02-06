package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(area string) ([]Pokemon, error){
	url := baseURL + "/location-area/" + area 

	data, present  := c.cache.Get(url)
	var areaData AreaData
	if !present{
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return make([]Pokemon, 0), fmt.Errorf("err creating request: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return make([]Pokemon, 0), fmt.Errorf("err making request: %w", err)
		}

		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return make([]Pokemon, 0), fmt.Errorf("err reading response: %w", err)
		}
		
		err = json.Unmarshal(dat,&areaData)
		if err != nil {
			return make([]Pokemon, 0), fmt.Errorf("err decoding response: %w", err)
		}
	}else{
		err := json.Unmarshal(data,&areaData)
		if err != nil {
			return make([]Pokemon, 0), fmt.Errorf("err decoding response %e", err)
		}
	}

	return parsePokemon(areaData), nil
}

func parsePokemon(a AreaData) []Pokemon{
	res := make([]Pokemon, len(a.PokemonEncounters))
	for i ,v := range a.PokemonEncounters{
		res[i] = v.Pokemon
	}
	return res
}