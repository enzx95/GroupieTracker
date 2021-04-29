package controller

import (
	"GroupieTracker/model"
	"encoding/json"
	"fmt"
)

func GetLocations() (*model.Locations, error) {

	body, err := MakeRequest("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	s, err := ParseLocations(body)
	return s, err
}

func ParseLocations(body []byte) (*model.Locations, error) {
	var locations = new(model.Locations)
	err := json.Unmarshal(body, locations)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return locations, err
}
