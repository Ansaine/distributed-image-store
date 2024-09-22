package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"strconv"
)

// decide which server has the the key
func getServer(key string) int {
	hash := fnv.New32a() // (Fowler-Noll-Vo) - a cnon cryptic hash function
	hash.Write([]byte(key))
	return int(hash.Sum32() % 3)
}

func Set(w http.ResponseWriter, r *http.Request) {

	// get data from post
	var requestData struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	fmt.Println("key " + requestData.Key + " and value " + requestData.Value + " to be sent")

	// calculate the server to send
	server := getServer(requestData.Key) + 1
	ip_end_point := server + 1

	// send the post data to server
	jsonData, _ := json.Marshal(requestData)
	url := "http://127.0.0." + strconv.Itoa(ip_end_point) + ":8080/set"
	println("url to route in set :", url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error sending POST request to db-server:", err)
		http.Error(w, "Error sending POST request to db-server", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK) // Set status code to 200
	w.Write([]byte("Image successfully stored")) // Send success message
}

// http://127.0.0.1:8080/get?key=hello
func Get(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	shard := getServer(key)
	server := shard + 1
	ip_end_point := server + 1
	url := "http://127.0.0." + strconv.Itoa(ip_end_point) + ":8080/get?key=" + key
	println("url to route in get:", url)

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Error sending the request to db-server : ", err)
	}

	// Read the response body and send it as it is
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Failed to read response from server", http.StatusInternalServerError)
		return
	}

	// Set the content type and write the response body back to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}
