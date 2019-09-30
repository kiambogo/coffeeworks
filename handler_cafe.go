package main

import (
	"io"
	"encoding/json"
	"net/http"
)

func GetCafes(w http.ResponseWriter, r *http.Request) {
	places, err := PlacesClient.FindPlacesNearArea()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}



	s := json.NewEncoder(w)
	s.SetIndent("", "    ")
	s.Encode(places)
}
