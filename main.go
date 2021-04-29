package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
	// "runtime"
)

type ArtistsApi struct {
	Id              int64    `json:"id"`
	Name            string   `json:"name"`
	Image           string   `json:"image"`
	Members         []string `json:"members"`
	CreationDate    uint16   `json:"creationDate"`
	LocationsUrl    string   `json:"locations"`
	ConcertDatesUrl string   `json:"concertDates"`
	RelationsUrl    string   `json:"relations"`
}

type ConcertDates struct {
	Id    int64    `json:"id"`
	Dates []string `json:"dates"`
}

type API struct {
}

func (r API) GetArtistsApi() (*[]ArtistsApi, error) {

	body, err := MakeRequest(`https://groupietrackers.herokuapp.com/api/artists`)
	if err != nil {
		return nil, err
	}
	s, err := ParseArtistsApi(body)
	return s, err
}

func (r API) GetConcertDates(id int64) (*ConcertDates, error) {

	body, err := MakeRequest("https://groupietrackers.herokuapp.com/api/dates/" + fmt.Sprint(id))
	if err != nil {
		return nil, err
	}
	s, err := ParseConcertDates(body)
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

func ParseArtistsApi(body []byte) (*[]ArtistsApi, error) {
	var artist = new([]ArtistsApi)
	err := json.Unmarshal(body, artist)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return artist, err
}

func ParseConcertDates(body []byte) (*ConcertDates, error) {
	var dates = new(ConcertDates)
	err := json.Unmarshal(body, dates)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return dates, err
}

func main() {
	fmt.Println("------")
	API := new(API)
	artists, err := API.GetArtistsApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println((*artists)[0])

	dates, err := API.GetConcertDates((*artists)[0].Id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dates)
}
