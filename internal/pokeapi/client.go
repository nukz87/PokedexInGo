package pokeapi

import (
	"net/http"
	"time"

	"github.com/nukz87/PokedexInGo/internal/pokecache"
)

type Client struct {
	pokePC     map[string]Pokemon
	pokeCache  pokecache.Cache
	httpClient http.Client
}

func NewClient(timeOut time.Duration) Client {
	return Client{
		pokeCache: pokecache.NewCache(5 * time.Second),
		httpClient: http.Client{
			Timeout: timeOut,
		},
	}
}
