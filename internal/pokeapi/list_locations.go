package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetLocationsList(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(resData, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}
