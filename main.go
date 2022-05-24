package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Spots struct {
	Spots []Spot `json:"spots"`
}

type Spot struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	// Coordinates int    `json:"coordinates"`
	// // struct {// Longitude float64 `json:"longitude"`
	// // 	// Latitude  float64 `json:"latitude"`} `json:"coordinates"`
	Link       string  `json:"link"`
	ImageName  string  `json:"imageName"`
	Difficulty float64 `json:"difficulty"`
	About      string  `json:"about"`
	Longitude  float64 `json: "longitude"`
	Latitude   float64 `json: "latitude"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the surf!")
}

func getAllSpots(w http.ResponseWriter, r *http.Request) {
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
	spotIDnum, _ := strconv.ParseInt(spotID, 10, 64)
	//sortir l'objet spot correspondant à ce spot id
	//refacto l'accès au json
	json.NewEncoder(w).Encode(spots.Spots[spotIDnum])
}

func addSpot(http.ResponseWriter, *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/floater")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	insert, err := db.Query("INSERT INTO `surf_spots` (`id`, `name`, `link`, `image`, `difficulty`, `about`, `longitude`, `latitude`) VALUES (NULL, 'pipeline', 'https://www.lonelyplanet.fr/place-be/surfer-le-banzai-pipeline', 'https://www.surf-forecast.com/system/images/4295/large/Banzai-Pipelines-and-Backdoor.jpg?1324521720', '3', '-158.05120469999997',)")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func getSpots(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/floater")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM `surf_spots`")

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()
	var spots []Spot
	for res.Next() {
		// byteValue, _ := ioutil.ReadAll(r.Body)

		var spot Spot
		err := res.Scan(&spot.Name, &spot.ID, &spot.Link, &spot.ImageName, &spot.Difficulty, &spot.About, &spot.Longitude, &spot.Latitude)
		spots = append(spots, spot)
		if err != nil {
			log.Fatal(err)
		}
		// json.Unmarshal(byteValue, &spots)
	}
	json.NewEncoder(w).Encode(spots)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", getAllSpots).Methods("GET")
	router.HandleFunc("/spots/{id}", getOneSpot).Methods("GET")
	router.HandleFunc("/addspots", addSpot).Methods("POST")
	router.HandleFunc("/getspots", getSpots).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
