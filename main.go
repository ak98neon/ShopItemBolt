package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const DefaultPort = ":8080"

func main() {
	Open()
	defer Close()

	r := mux.NewRouter()
	r.HandleFunc("/items", GetAllItems).Methods("GET")
	r.HandleFunc("/items/{id}", GetItemById).Methods("GET")
	r.HandleFunc("/items", SaveItem).Methods("POST")
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(DefaultPort, handlers.CORS(corsObj)(r)))
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	item, err := GetItem(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	_ = json.NewEncoder(w).Encode(item)
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()
	size, errSize := strconv.Atoi(params.Get("size"))
	page, errPage := strconv.Atoi(params.Get("page"))

	items := List(ItemBucket)
	if errSize != nil || errPage != nil {
		pageRequest := PageRequest{
			Items:       items,
			TotalCount:  len(items),
			CountOfPage: len(items),
			CurrentPage: 1,
		}
		_ = json.NewEncoder(w).Encode(pageRequest)
	} else {
		log.Println("Request size: ", size, " and request page: ", page)
		start := (page - 1) * size
		end := start + size

		if start < 0 || start > len(items) {
			http.Error(w, "Bad request", 400)
		}

		if end > len(items) {
			end = len(items) - 1
		}

		paginationItems := items[start:end]
		pageRequest := PageRequest{
			Items:       paginationItems,
			TotalCount:  len(paginationItems),
			CountOfPage: size,
			CurrentPage: page,
		}
		_ = json.NewEncoder(w).Encode(pageRequest)
	}
}

func SaveItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.GenerateUniqueId()
	_ = item.Save()
	_ = json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)
	id := params["id"]
	item.ID = id

	_ = item.Save()
	_ = json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		log.Fatal("Id is empty!")
	}
	_ = Delete(id)
	_ = json.NewEncoder(w).Encode("Item was deleted!")
}
