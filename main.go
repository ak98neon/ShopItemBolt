package main

import (
	db "akudria/appleShop/pudge"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const DEFAULT_PORT = ":8080"
const ItemBucket = "item"

func main() {
	db.Open()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/items", GetAllItems).Methods("GET")
	r.HandleFunc("/items", SaveItem).Methods("POST")
	log.Fatal(http.ListenAndServe(DEFAULT_PORT, r))
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.List(ItemBucket)
	_ = json.NewEncoder(w).Encode(items)
}

func SaveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item db.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.GenerateUniqueId()
	_ = item.Save()
	_ = json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {

}
