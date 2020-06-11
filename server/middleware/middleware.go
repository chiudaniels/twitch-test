package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../model"

	"net/http"
)

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var NewMovie model.Entry
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &NewMovie)
	fmt.Println(NewMovie)
	// _ = json.NewDecoder(r.Body).Decode(&NewMovie)
	// fmt.Println(json.NewDecoder(r.Body))
	json.NewEncoder(w).Encode(NewMovie)
	model.AddEntry(NewMovie.Year, NewMovie.Title)
}

func GetAllEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	entries := model.GetAllEntry()
	json.NewEncoder(w).Encode(entries)
	fmt.Println(entries)
}
