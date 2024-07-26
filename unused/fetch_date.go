package Fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"main.go/Errors"
)

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func Fetch_date(w http.ResponseWriter, r *http.Request, url string) Date {
	var datesURL = url

	// Fetch the data from the URL
	response, err := http.Get(datesURL)
	if err != nil {
		Errors.Error500(w, r)
	}
	defer response.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Errors.Error500(w, r)
	}

	// Unmarshal the JSON data
	var jsonDates Date
	err = json.Unmarshal(data, &jsonDates)
	if err != nil {
		Errors.Error500(w, r)
	}

	
	//Removing the * from the dates
	for i := 0; i < len(jsonDates.Dates); i++ {
		jsonDates.Dates[i] = jsonDates.Dates[i][1:]

	}
	return jsonDates
}
