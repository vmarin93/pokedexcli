package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	if err := json.Unmarshal(resData, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
