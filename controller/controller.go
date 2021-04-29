package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

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

func Init() {

	start := time.Now()
	fmt.Printf("\n")
	artists, err := GetArtistsApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println((*artists)[0])

	dates, err := GetConcertDates()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dates.Index[0].Dates)

	locations, err := GetLocations()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(locations.Index[0].Locations)

	relation, err := GetRelation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(relation.Index[0])

	elapsed := time.Since(start)
	fmt.Printf("\ntook %s \n", elapsed)
}
