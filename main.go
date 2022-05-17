package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gorilla/mux"
)

type Spots struct {
    Spots []Spot `json:"spots"`
}

type Spot struct {
	Name        string `json:"name"`
	ID          int    `json:"id"`
	Coordinates struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	} `json:"coordinates"`
	Link       string `json:"link"`
	ImageName  string `json:"imageName"`
	Difficulty int    `json:"difficulty"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the surf!")
}

func printJson(w http.ResponseWriter, r *http.Request) {
	// Open our jsonFile
	jsonFile, err := os.Open("spots.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened spots.json")
	
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	
	// we initialize our Users array
	var spots Spots
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &spots)

	for i := 0; i < len(spots.Spots); i++ {
		fmt.Println("Spot Name: " + spots.Spots[i].Name)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", printJson)
	log.Fatal(http.ListenAndServe(":8080", router))
}