package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Define a structure that matches the expected JSON format
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Make the HTTP GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal("Error fetching the data:", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// Unmarshal the response body into the Todo struct
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatal("Error unmarshalling response:", err)
	}

	// Print the response data
	fmt.Printf("Fetched Todo:\nUserID: %d\nID: %d\nTitle: %s\nCompleted: %v\n", todo.UserID, todo.ID, todo.Title, todo.Completed)
}
