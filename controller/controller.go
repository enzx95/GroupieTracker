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

func artistsInit() {
	artists, err := GetArtistsApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Artists done.")
	fmt.Println((*artists)[0])
}

func datesInit() {
	dates, err := GetConcertDates()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Dates done.")
	fmt.Println(dates.Index[0].Dates)
}
func locationsInit() {
	locations, err := GetLocations()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Locations done.")
	fmt.Println(locations.Index[0].Locations)
}
func relationInit() {
	relation, err := GetRelation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Relation done.")
	fmt.Println(relation.Index[0])
}

func Init() {

	start := time.Now()
	fmt.Printf("\n")

	go artistsInit()
	go datesInit()
	go locationsInit()
	go relationInit()

	time.Sleep(600 * time.Millisecond)
	// from 800ms to 600ms
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s \n", elapsed)

}
