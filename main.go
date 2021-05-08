package main

import (
	"GroupieTracker/controller"
	"GroupieTracker/model"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

var tableau [52]*model.Artist

func makeTab() {
	start := time.Now()
	for i := 0; i < 52; i++ {
		tableau[i] = controller.GetDataByID(i)
	}
	elapsed := time.Since(start)
	fmt.Printf("\ntook %s \n", elapsed)

}

func join(s ...string) string {
	// first arg is sep, remaining args are strings to join
	return strings.Join(s[1:], s[0])
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("Artists2.html").Funcs(template.FuncMap{"join": join}).ParseFiles("./assets/pages/Artists2.html")

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, tableau); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	controller.Init()
	makeTab()
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", mainPageHandler)
	http.ListenAndServe(":8080", nil)

}
