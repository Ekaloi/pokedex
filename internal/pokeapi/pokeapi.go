package pokeapi

import (
	"net/http"
	"time"

	"github.com/Ekaloi/pokedex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)


// Client -
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{},
		cache: pokecache.NewCache(5 * time.Second),
	}
}