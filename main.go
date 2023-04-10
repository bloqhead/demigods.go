package main

import (
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

type Scaling struct {
	Str string `json:"str,omitempty"`
	Int string `json:"int,omitempty"`
	Fai string `json:"fai,omitempty"`
	Arc string `json:"arc,omitempty"`
	Dex string `json:"dex,omitempty"`
}

type Stat struct {
	Weight float64 `json:"weight,omitempty"`
	Physical float64 `json:"physical,omitempty"`
	Magic float64 `json:"magic,omitempty"`
	Fire float64 `json:"fire,omitempty"`
	Light float64 `json:"light,omitempty"`
	Holy float64 `json:"holy,omitempty"`
}

type Weapon struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Tier string `json:"tier"`
	Category string `json:"category"`
	Scaling Scaling `json:"scaling"`
	Skill string `json:"skill"`
	Stats Stat `json:"stats"`
}

var weapons []Weapon

// data setup
func setup() {
	content, err := ioutil.ReadFile("./data/all.json")

	if err != nil {
		log.Fatal("Error opening data file: ", err)
	}

	err = json.Unmarshal(content, &weapons)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}

// fetch all items
func fetchAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weapons)
}

// by ID
func fetchById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	items := make([]Weapon, 0, len(weapons))
	
	for _, v := range weapons {
		if strconv.Itoa(v.ID) == id {
			items = append(items, v)
		}
	}

	json.NewEncoder(w).Encode(items)
}

// by category
func fetchByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	category := vars["category"]
	items := make([]Weapon, 0, len(weapons))
	
	for _, v := range weapons {
		cat := slug.Make(v.Category)
		slug := slug.Make(category)

		if cat == slug {
			items = append(items, v)
		}
	}

	json.NewEncoder(w).Encode(items)
}

// category list
func fetchCategoryList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// request handler
func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(mux.CORSMethodMiddleware(r))
	
	setup()

	// home
	// r.HandleFunc("/", homePage)
	
	// all items
	r.HandleFunc("/all", fetchAllItems)

	// by id
	r.HandleFunc("/{id}", fetchById)
	
	// by category
	r.HandleFunc("/category/{category}", fetchByCategory)

	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	handleRequests()
}
