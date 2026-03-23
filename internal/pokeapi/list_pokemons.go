package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemonsAtLocation(location_area string) (Pokemons, error) {
	url := baseURL + "/location-area/" + location_area

	cacheData, ok := client.cache.Get(url)
	if ok {
		pokemons := Pokemons{}
		if err := json.Unmarshal(cacheData, &pokemons); err != nil {
			return Pokemons{}, err
		}
		return pokemons, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemons{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemons{}, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemons{}, err
	}
	client.cache.Add(url, resData)

	pokemons := Pokemons{}
	if err := json.Unmarshal(resData, &pokemons); err != nil {
		return Pokemons{}, err
	}
	return pokemons, nil
}
