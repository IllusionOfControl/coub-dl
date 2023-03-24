package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchHttp(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}

func FetchCoubMetadata(coubId string) (map[string]interface{}, error) {
	apiUrl := fmt.Sprintf("https://coub.com/api/v2/coubs/%s", coubId)

	data, err := fetchHttp(apiUrl)
	if err != nil {
		return nil, err
	}
	metadata := map[string]interface{}{}
	if err = json.Unmarshal(data, &metadata); err != nil {
		log.Fatalln(err)
	}
	return metadata, nil
}

func FetchCoubFile(url string) ([]byte, error) {
	return fetchHttp(url)
}
