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

	if r.URL.Path != "/" || r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	tmpl, err := template.New("Artists.html").Funcs(template.FuncMap{"join": join}).ParseFiles("./assets/pages/Artists.html")

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, tableau); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("./assets/pages/error.html")
		if err != nil {
			log.Fatal(err.Error())
		}

		err = t.Execute(w, "404 Not Found")
		if err != nil {
			log.Fatal(err.Error())
		}
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
