package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"F:/assets/model"
	"F:/assets/controller"

)



func handleRequest() {
	new := mux.NewRouter().StrictSlash(true)
	new.HandleFunc("/", homepage)
	new.HandleFunc("/returnAll", returnAllAssets)
	new.HandleFunc("/create", createNewAsset).Methods("POST")
	new.HandleFunc("/update/{id}", updateAsset).Methods("PUT")
	new.HandleFunc("/delete/{id}", deleteAsset).Methods("DELETE")
	new.HandleFunc("/return/{id}", returnSingleAsset)
	log.Fatal(http.ListenAndServe(":10000", new))
}

//main function which will kick off our API.
func main() {
	fmt.Println("crypto")
	Assets = []Asset{
		Asset{Id: "1", AssetName: "XRP" },
		Asset{Id: "2", AssetName: "Bitcoin"},
	}
	handleRequest()
}
