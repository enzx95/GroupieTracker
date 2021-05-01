package main

import (
	"GroupieTracker/controller"
	"GroupieTracker/model"
	"log"
	"net/http"
	"text/template"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

	var tableau [3]*model.Artist
	tableau[0] = controller.GetDataByID(0)
	tableau[1] = controller.GetDataByID(1)
	tableau[2] = controller.GetDataByID(2)

	if err := tmpl.Execute(w, tableau); err != nil {
		log.Fatal(err)
	}

	//tmpl.Execute(w, "Karim")

}

func main() {
	controller.Init()
	controller.GetDataByID(0)

	http.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":8080", nil)

}
