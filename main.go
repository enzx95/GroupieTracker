package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"GroupieTracker/model"
)

type API struct {
}

func (r API) GetArtistsApi() (*[]model.ArtistsApi, error) {

	body, err := MakeRequest(`https://groupietrackers.herokuapp.com/api/artists`)
	if err != nil {
		return nil, err
	}
	s, err := ParseArtistsApi(body)
	return s, err
}

func (r API) GetConcertDates() (*model.Dates, error) {

	body, err := MakeRequest("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	s, err := ParseConcertDates(body)
	return s, err
}

func (r API) GetLocations() (*model.Locations, error) {

	body, err := MakeRequest("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	s, err := ParseLocations(body)
	return s, err
}

func (r API) GetRelation() (*model.Relation, error) {

	body, err := MakeRequest("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	s, err := ParseRelation(body)
	return s, err
}

func MakeRequest(url string) ([]byte, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), err
}

func ParseArtistsApi(body []byte) (*[]model.ArtistsApi, error) {
	var artist = new([]model.ArtistsApi)
	err := json.Unmarshal(body, artist)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return artist, err
}

func ParseConcertDates(body []byte) (*model.Dates, error) {
	var dates = new(model.Dates)
	err := json.Unmarshal(body, dates)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return dates, err
}

func ParseLocations(body []byte) (*model.Locations, error) {
	var locations = new(model.Locations)
	err := json.Unmarshal(body, locations)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return locations, err
}

func ParseRelation(body []byte) (*model.Relation, error) {
	var relation = new(model.Relation)
	err := json.Unmarshal(body, relation)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return relation, err
}

func main() {
	//fmt.Println("------")
	API := new(API)
	// artists, err := API.GetArtistsApi()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println((*artists)[0])

	// dates, err := API.GetConcertDates()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(dates.Index[0].Dates)

	// locations, err := API.GetLocations()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(locations.Index[0].Locations)

	relation, err := API.GetRelation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(relation.Index[0])
}
