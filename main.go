package main

import (
	db "akudria/appleShop/pudge"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const DefaultPort = ":8080"

func main() {
	db.Open()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/items", GetAllItems).Methods("GET")
	r.HandleFunc("/items/{id}", GetItemById).Methods("GET")
	r.HandleFunc("/items", SaveItem).Methods("POST")
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(DefaultPort, r))
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item, err := db.GetItem(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	_ = json.NewEncoder(w).Encode(item)
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	items := db.List(db.ItemBucket)
	_ = json.NewEncoder(w).Encode(items)
}

func SaveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var item db.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.GenerateUniqueId()
	_ = item.Save()
	_ = json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	var item db.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)
	id := params["id"]
	item.ID = id

	_ = item.Save()
	_ = json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		log.Fatal("Id is empty!")
	}
	_ = db.Delete(id)
	_ = json.NewEncoder(w).Encode("Item was deleted!")
}
