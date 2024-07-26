package Fetch

import (
	"fmt"
	"net/http"
	"strconv"
)

func Fetch_profile(w http.ResponseWriter, r *http.Request, jsonArtistsCards []Artist) *Artist {
	// Parse artist ID from URL query parameter
	artistID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid artist ID: %v", err)
		return nil
	}
	var data_location [][]string
	data_location = Fetch_Relation(w, r, jsonArtistsCards[artistID].Relations)

	// Find the artist in jsonArtistsCards
	var artistInfo *Artist
	for _, artist := range jsonArtistsCards {
		if artist.ID == artistID {
			artistInfo = &artist
			break
		}
	}
	artistInfo.Date_Locat = append(artistInfo.Date_Locat, data_location...)

	fmt.Print(artistInfo.Date_Locat)
	// Handle case where artist is not found
	if artistInfo == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Artist Not Found")
		return nil
	}

	// Return the artist information
	return artistInfo
}
