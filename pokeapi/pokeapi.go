package pokeapi


import (
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)


// Client -
type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
	}
}