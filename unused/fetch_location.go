package Fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"main.go/Errors"
)

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

func Fetch_location(w http.ResponseWriter, r *http.Request, url string) Location {
	response, err := http.Get(url)
	if err != nil {
		Errors.Error500(w, r)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Errors.Error500(w, r)
	}

	var jsonLocations Location

	err = json.Unmarshal(data, &jsonLocations)

	if err != nil {
		Errors.Error500(w, r)
	}

	for _, i := range jsonLocations.Locations {
		fmt.Print(i)

	}

	return jsonLocations

}
