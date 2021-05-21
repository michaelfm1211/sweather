package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

var templates *template.Template = template.Must(template.ParseGlob("templates/*.html"))

func weather(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	coorinates, err := geocode(vars["addr"])
	if err != nil {
		templates.ExecuteTemplate(w, "error.html", err)
		return
	}

	forecast, err := Forecasts(coorinates)
	if err != nil {
		templates.ExecuteTemplate(w, "error.html", err)
		return
	}

	templates.ExecuteTemplate(w, "weather.html", forecast)
}

func index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates.ExecuteTemplate(w, "index.html", nil)
	case http.MethodPost:
		r.ParseForm()
		addr := r.PostForm.Get("addr")
		http.Redirect(w, r, fmt.Sprintf("/weather/%s", addr), http.StatusFound)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT not defined")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET", "POST")
	r.HandleFunc("/weather/{addr}", weather).Methods("GET")

	staticServer := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", staticServer))

	log.Printf("Server up on port %s\n", port)
	log.Fatal(http.ListenAndServe(":" + port, r))
}
