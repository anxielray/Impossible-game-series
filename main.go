package main

import (
	m "Zuxyll-game/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	//serve the home page
	http.HandleFunc("/", m.HomePage)

	//login
	//register
	//server the game page
	http.HandleFunc("/play", m.Play)

	//
	http.HandleFunc("/win", m.Win)

	http.HandleFunc("/loose", m.Loose)

	//serve the game
	http.HandleFunc("/step", m.Game)
	/* the game terminates when you make a trip and land a step where you were earlier*/

	//define the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "1919"
	}

	//start the server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
