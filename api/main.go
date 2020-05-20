package main

import (
	// "encoding/json"
	// "io/ioutil"
	"log"
	"net/http"
	"os"

	//	"time"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func main() {
	// Connect to mongo
	session, err := mgo.Dial("mongo:27017")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("mongo err")
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Set up routes
	r := mux.NewRouter()

	http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	log.Println("Listening on port 8080...")
}
