package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"` // Exported, so JSON can use it
}

// In-memory storage for storing users
var userCache = make(map[int]User)

var cacheMutex sync.RWMutex // Mutex for making thread safe

func main() {
	mux := http.NewServeMux() // Handler to handle incoming requests

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /user", createUser)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server Listening to: 8080")
	http.ListenAndServe(":8080", mux)

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()

	if !ok {
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	if user.Name == "" {
		http.Error(w,
			"Name is required",
			http.StatusBadRequest,
		)

		return
	}

	cacheMutex.Lock()
	// Store the user in the in-memory user cache
	userCache[len(userCache)+1] = user

	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	_, ok := userCache[id]

	if !ok {
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	cacheMutex.Lock()
	delete(userCache, id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
