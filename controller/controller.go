package controller

import (
	"GroupieTracker/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var artist *[]model.ArtistsApi
var dates *model.Dates
var locations *model.Locations
var relation *model.Relation

var colorGreen = "\033[32m"
var colorReset = "\033[0m"

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
	artistsData, err := GetArtistsApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(colorGreen), "Artists done.")
	fmt.Print(string(colorReset))
	fmt.Println((*artistsData)[0])
	artist = artistsData
}
func datesInit() {
	datesData, err := GetConcertDates()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(colorGreen), "Dates done.")
	fmt.Print(string(colorReset))
	fmt.Println(datesData.Index[0].Dates)
	dates = datesData
}
func locationsInit() {
	locationsData, err := GetLocations()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(colorGreen), "Locations done.")
	fmt.Print(string(colorReset))
	fmt.Println(locationsData.Index[0].Locations)
	locations = locationsData
}
func relationInit() {
	relationData, err := GetRelation()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(colorGreen), "Relation done.")
	fmt.Print(string(colorReset))
	fmt.Println(relationData.Index[0])
	relation = relationData
}

func GetDataByID(id int) *model.Artist {

	var artistData = new(model.Artist)
	artistData.Id = int64(id)
	artistData.Name = (*artist)[id].Name
	artistData.Image = (*artist)[id].Image
	artistData.Members = (*artist)[id].Members

	for i, s := range dates.Index[id].Dates {
		dates.Index[id].Dates[i] = s[1:]
	}

	artistData.ConcertDates = dates.Index[id].Dates

	for i, s := range locations.Index[id].Locations {
		locations.Index[id].Locations[i] = strings.ReplaceAll(s, "_", " ")
	}

	artistData.ConcertLocations = locations.Index[id].Locations
	artistData.CreationDate = (*artist)[id].CreationDate
	artistData.Relation = relation.Index[id].DatesLocations

	empJSON, err := json.MarshalIndent(artistData, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf(string(empJSON))
	return artistData

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
