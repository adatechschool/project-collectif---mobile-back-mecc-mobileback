package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

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
	Difficulty float64 `json:"difficulty"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the surf!")
}

func getAllSpots(w http.ResponseWriter, r *http.Request){
	jsonFile, err := os.Open("spots.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened spots.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var spots Spots
	json.Unmarshal(byteValue, &spots)
	json.NewEncoder(w).Encode(spots.Spots)
}

func getOneSpot(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("spots.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened spots.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var spots Spots
	json.Unmarshal(byteValue, &spots)
	spotID := mux.Vars(r)["id"]
	spotIDnum,_ := strconv.ParseInt(spotID, 10, 64)
	//sortir l'objet spot correspondant à ce spot id
	//refacto l'accès au json
	json.NewEncoder(w).Encode(spots.Spots[spotIDnum])
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", getAllSpots).Methods("GET")
	router.HandleFunc("/spots/{id}", getOneSpot).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}