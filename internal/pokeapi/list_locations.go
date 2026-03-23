package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (client *Client) GetLocationsList(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cacheData, ok := client.cache.Get(url)
	if ok {
		locations := Locations{}
		err := json.Unmarshal(cacheData, &locations)
		if err != nil {
			return Locations{}, err
		}
		log.Println("Retrieved from cache")
		return locations, nil
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

	client.cache.Add(url, resData)

	locations := Locations{}
	err = json.Unmarshal(resData, &locations)
	if err != nil {
		return Locations{}, err
	}

	log.Println("Retrieved from network")
	return locations, nil
}
