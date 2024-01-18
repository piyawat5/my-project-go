package main

import (
	"log"
	"net/http"

	"github.com/piyawat5/myproject/routes/movie"
)

func main() {
	http.HandleFunc("/movies",movie.MoivesHandler)

	err:=http.ListenAndServe("localhost:3100",nil)
	log.Fatal(err)
}