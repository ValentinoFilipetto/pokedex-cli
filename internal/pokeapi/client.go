package pokeapi

import (
	"net/http"
	"time"

	"github.com/ValentinoFilipetto/pokedex-cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(timeout),
	}
}
