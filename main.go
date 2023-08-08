package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID         string     `json:"id"`
	SecretCode string     `json:"secretCode"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Playlists  []Playlist `json:"complaints"`
}

type Playlist struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Songs []Song `json:"songs"`
}

type Song struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Composers string `json:"composers"`
	MusicURL  string `json:"music_url"`
}

var usersDB = make(map[string]User)

func ReturnJsonResponse(res http.ResponseWriter, resMessage []byte) {
	res.Header().Set("content-type", "application/json")
	res.Write(resMessage)
}

func generateUniqueID() string {
	rand.Seed(time.Now().UnixNano())

	uniqueID := strconv.Itoa((rand.Intn(10000)))

	return uniqueID
}

func generateUniqueSecretCode() string {
	// Implement a logic to generate a unique secret code
	rand.Seed(time.Now().UnixNano())

	secretCode := strconv.Itoa((rand.Intn(1000000)))

	return secretCode
}

// Implement the handler functions for each route

func loginUser(secretCode string) (User, error) {
	// Loop through the usersDB to find a user with the provided secret code
	for _, user := range usersDB {
		if user.SecretCode == secretCode {
			return user, nil
		}
	}

	// If no user with the provided secret code is found, return an error
	return User{}, fmt.Errorf("invalid secret code")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /login route
	if r.Method != "POST" {
		HandlerMessage := []byte(`{
			"success" : false,
			"message" :"check your HTTP method : Invalid HTTP method executed",
		}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// Parse the request body to get the user's secret code
	var requestBody struct {
		SecretCode string `json:"secretCode"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		HandlerMessage := []byte(`{
			"success" : false,
			"message" : "Error parsing the req body data",
		}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// Call the loginUser function to authenticate the user
	user, err := loginUser(requestBody.SecretCode)

	if err != nil {
		HandlerMessage := []byte(`{
			"success":false,
			"message":"Wrong secret code/user not found",
 }`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// Convert the user to JSON format
	userJSON, err := json.Marshal(user)

	if err != nil {
		HandlerMessage := []byte(`{
   "success":false,
   "message":"Error parsing the user data",
}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// Set the Content-Type header and respond with the user details
	HandlerMessage := []byte(`{
		"success" : true,
		"message" : "user sign-in successfully",
	}`)
	ReturnJsonResponse(w, HandlerMessage)
	ReturnJsonResponse(w, userJSON)
	return
}
func registerUser(name, email string) User {
	userID := generateUniqueID()             // Implement a function to generate a unique ID
	secretCode := generateUniqueSecretCode() // Implement a function to generate a unique secret code
	newUser := User{
		ID:         userID,
		SecretCode: secretCode,
		Name:       name,
		Email:      email,
		Playlists:  []Playlist{},
	}
	usersDB[userID] = newUser
	return newUser
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /register route
	if r.Method != "POST" {
		HandlerMessage := []byte(`{
			"success" : false,
			"message" :"check your HTTP method : Invalid HTTP method executed",
		}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// Parse the request body to get the name and email of the new user
	var newUser struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		HandlerMessage := []byte(`{
			"success" : false,
			"message" : "Error parsing the request body data",
		}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}
	_, ok := usersDB[newUser.Email]
	if ok {
		HandlerMessage := []byte(`{
	 		"success" : false,
	 		"message" : "User already exist",
	 	}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}
	// Call the registerUser function to create a new user
	user := registerUser(newUser.Name, newUser.Email)

	// Convert the user to JSON format
	userJSON, err := json.Marshal(user)

	if err != nil {
		HandlerMessage := []byte(`{
	       "success":false,
	       "message":"Error parsing the  data",
}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// Set the Content-Type header and respond with the user details
	HandlerMessage := []byte(`{
		"success" : true,
		"message" : "New user sign-up",
	}`)
	ReturnJsonResponse(w, HandlerMessage)
	ReturnJsonResponse(w, userJSON)
	return

}

func viewProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /viewProfile route

	if r.Method != "GET" {
		HandlerMessage := []byte(`{
			"success":false,
            "message":"check your HTTP method : Invalid HTTP method executed",
		}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	userID := "34567" // Example: Extract user ID from headers

	// Find the user in the usersDB using the userID
	user, found := usersDB[userID]
	if !found {
		HandlerMessage := []byte(`{
			"success":false,
            "message":"User not found",
		}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// Return the user's playlists as JSON response
	responseJSON, err := json.Marshal(user.Playlists)

	if err != nil {
		HandlerMessage := []byte(`{
			"success":false,
			"message":"Error parsing the  data",
 }`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	ReturnJsonResponse(w, responseJSON)
}

func getAllSongsOfPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /getAllSongsOfPlaylist route
	if r.Method != "GET" {
		HandlerMessage := []byte(`{
			"success":false,
            "message":"check your HTTP method : Invalid HTTP method executed",
		}`)
		ReturnJsonResponse(w, HandlerMessage)
		return
	}

	// playlistID := "7899" // Example: ?playlistId=your_playlist_id

	// Find the playlist by ID from a hypothetical playlist database
	// playlist, found := usersDB[playlistID]
	// if !found {
	// 	http.Error(w, "Playlist not found", http.StatusNotFound)
	// 	return
	// }

	// Return playlist's songs as JSON response
	// responseJSON, err := json.Marshal(playlist.Songs)
	// if err != nil {
	// 	http.Error(w, "Error encoding response", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(responseJSON)
}

func createPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /createPlaylist route
}

func addSongToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /addSongToPlaylist route
}

func deleteSongFromPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /deleteSongFromPlaylist route
}

func deletePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /deletePlaylist route
}

func getSongDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for /getSongDetail route
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/viewProfile", viewProfileHandler)
	http.HandleFunc("/getAllSongsOfPlaylist", getAllSongsOfPlaylistHandler)
	http.HandleFunc("/createPlaylist", createPlaylistHandler)
	http.HandleFunc("/addSongToPlaylist", addSongToPlaylistHandler)
	http.HandleFunc("/deleteSongFromPlaylist", deleteSongFromPlaylistHandler)
	http.HandleFunc("/deletePlaylist", deletePlaylistHandler)
	http.HandleFunc("/getSongDetail", getSongDetailHandler)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
