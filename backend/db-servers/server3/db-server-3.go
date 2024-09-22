package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// structure of a store
type Store struct {
	data map[string]string
	file string
}

// using receiver to define load method of struct
// data from file is loaded onto slice byte and then unmarshalled onto store's data
func (s *Store) load() {

	if _, err := os.Stat(s.file); os.IsNotExist(err) {

		// create file for store if not present
		file, err := os.Create(s.file)
		if err != nil {
			fmt.Println("error creating new file for store : ", err)
			return
		}
		defer file.Close()
	} else {
		// read file if present
		bytes, err := os.ReadFile(s.file)
		if err != nil {
			fmt.Println("Error reading file", err)
		} else {
			json.Unmarshal(bytes, &s.data)
		}
	}
}

func (s *Store) save() {
	bytes, err := json.Marshal(s.data)
	if err != nil {
		fmt.Println("Error marshelling data : ", err)
		return
	} else {
		err = os.WriteFile(s.file, bytes, 0644)
		if err != nil {
			fmt.Println("Error writing bytes into file : ", err)
		}
	}
}

// function to create a db
func newStore(file string) *Store {
	store := &Store{file: file, data: make(map[string]string)}
	store.load()
	return store
}

func (s *Store) get(key string) (string, bool) {
	// map returns value as well as if it exists
	value, exists := s.data[key]
	return value, exists
}

func (s *Store) set(key string, value string) {
	s.data[key] = value
	s.save()
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing 'key' query parameter", http.StatusBadRequest)
		return
	}

	db3 := newStore("db3.json")
	value, exists := db3.get(key)
	if !exists {
		value = "Does not exist"
	}

	response := map[string]string{
		"key":   key,
		"value": value,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func setHandler(w http.ResponseWriter, r *http.Request) {

	// get the request data
	var requestData struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// set the key value in store
	db3 := newStore("db3.json")
	db3.set(requestData.Key, requestData.Value)

	// Send back the response as JSON
	response := map[string]string{
		"message": "key set in db 3",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	port := ":8080"
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)

	fmt.Println("db-server-3 starting on 127.0.0.4", port)
	err := http.ListenAndServe("127.0.0.4"+port, nil)
	if err != nil {
		fmt.Println("error starting db-server-3 : ", err)
	}
}
