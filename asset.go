package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//asset structure
type Asset struct {
	AssetName     string  `json:"AssetName"`
	AssetType     string  `json:"AssetType"`
	AssetCode     string  `json:"AssetCode"`
	Id            int     `json:"Id"`
	CurrentPrice  float64 `json:"CurrentPrice"`
	PreviousPrice float64 `json:"PreviousPrice"`
}

//global Assets array that we can then populate in our main function  to simulate a database
var Assets []Asset

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func returnAllAssets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllAssets")
	json.NewEncoder(w).Encode(Assets)
}

func returnSingleAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, asset := range Assets {
		if asset.Id == key {
			json.NewEncoder(w).Encode(asset)
		}
	}
}

func createNewAsset(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var asset Asset
	json.Unmarshal(requestBody, &asset)
	Assets = append(Assets, asset)
	json.NewEncoder(w).Encode(asset)

}

func deleteAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for index, asset := range Assets {
		if asset.Id == key {
			Assets = append(Assets[:index], Assets[index+1:]...)
		}
	}

}

func updateAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var updatedEvent Asset
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedEvent)
	for i, asset := range Assets {
		if asset.Id == key {
			asset.Title = updatedEvent.Title
			asset.Desc = updatedEvent.Desc
			asset.Content = updatedEvent.Content
			Assets[i] = asset
			json.NewEncoder(w).Encode(asset)
		}
	}
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/assets", returnAllAssets)
	myRouter.HandleFunc("/asset", createNewAsset).Methods("POST")
	myRouter.HandleFunc("/asset/{id}", updateAsset).Methods("PUT")
	myRouter.HandleFunc("/asset/{id}", deleteAsset).Methods("DELETE")
	myRouter.HandleFunc("/asset/{id}", returnSingleAsset)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

//main function which will kick off our API.
func main() {
	fmt.Println("crypto")
	handleRequest()
}
