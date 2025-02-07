package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations , error){
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	locationsResp := RespShallowLocations{}
	data, present := c.Cache.Get(url)

	if !present{
			req, err := http.NewRequest("GET", url, nil)
		if err != nil{
			return RespShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()

		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.Cache.Add(url,dat)
		
		err = json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
	}else{
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
	}

	

	

	return locationsResp, nil

}