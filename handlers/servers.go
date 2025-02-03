package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").ParseFiles("public/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}


func Loose(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("loose.html").ParseFiles("public/templates/loose.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Win(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("win.html").ParseFiles("public/templates/win.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Play(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").ParseFiles("public/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
