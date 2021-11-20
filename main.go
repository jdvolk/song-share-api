package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"io/ioutil"

	"github.com/gorilla/mux"
)

//Data Arrays/Globals 

// SearchResults is ...
var SearchResults []SongDetails

// TimelinePosts is ...
var TimelinePosts []SongPost

// UserFavorites is ...
var UserFavorites []SongDetails

// UserData is ...
var UserData UserDetails



// Homepage of api
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	// w.Write([]byte("/"))
}

// Request Handler
func handleRequests() {
	myRouter := mux.NewRouter()
	// myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/searchResults",  func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "localhost:10000")
	// }).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	// myRouter.Use(mux.CORSMethodMiddleware(myRouter))
	myRouter.HandleFunc("/Home", returnAllTimeline)
	myRouter.HandleFunc("/Home/{ID}", addComment).Methods("POST")
	// myRouter.HandleFunc("/Home/{ID}", returnSingleTimeline)
	myRouter.HandleFunc("/searchResults", createNewSongPost).Methods("POST")
	myRouter.HandleFunc("/searchResults", returnSearchResults)
	myRouter.HandleFunc("/searchResults/{ID}", returnSingleSearchResult)
	myRouter.HandleFunc("/search/{SearchTerm}", searchItunesForArtistID)
	myRouter.HandleFunc("/User", returnUser)
	myRouter.HandleFunc("/favorites", addFavorite).Methods("POST")
	myRouter.HandleFunc("/favorites", returnUserFavorites)
	myRouter.HandleFunc("/favorites/{ID}", deleteFavorite).Methods("OPTIONS","DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

//Search Results Requests



func searchItunesForArtistID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchTerm := vars["SearchTerm"]
	url1 := "https://itunes.apple.com/search?term="
	url2 := "&country=US&entity=song,album,podcast"
	// var fullURL = `https://itunes.apple.com/search?term=btycll&country=US&entity=song,album,podcast`

	// var url1 = "https://itunes.apple.com/search?term="
	// var url2 = "&country=US&entity=song,album,podcast"
	var fullURL = url1 + searchTerm + url2 
	resp, err := http.Get(fullURL)
	if err != nil {
  log.Fatalln(err)
	}
	// body, err := ioutil.ReadAll(resp.Body)
	var data map[string]interface {}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}
//Convert the body to type string
// sb := string(body)
	
	log.Printf("%+v\n", data)
	log.Printf("%+v\n", data["results"].([]interface{})[0])

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
	return
}



func returnSearchResults(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnSearchResults")
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(SearchResults)
}

func returnSingleSearchResult(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["ID"]
	keyInt, _ := strconv.Atoi(key)
	// TODO make a new function that takes the key  and *where to search* and returns an error or a number
	// so this function isnt doing all of that
	for _, song := range SearchResults {
		if song.SongID == (keyInt) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(song)
		}
	}
}

// Timeline/Home Requests
func returnAllTimeline(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnTimelinePosts")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(TimelinePosts)
}

func returnSingleTimeline(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["ID"]
	keyInt, _ := strconv.Atoi(key)
	// TODO make a new function that takes the key and *where to search* and returns an error or a number
	// so this function isnt doing all of that
	for _, post := range TimelinePosts {
		if post.PostID == (keyInt) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(post)
		}
	}
}

func createNewSongPost(w http.ResponseWriter, r *http.Request){
	fmt.Println("POST Endpoint Hit: createSongPost")
	// get the body of our POST request
	// unmarshal this into a new Song struct
	// append this to our SongResults array.    
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(reqBody), &reqBody)
	var post SongPost 
	json.Unmarshal(reqBody, &post)
	// update our global SongPost array to include
	// our new post
	TimelinePosts = append(TimelinePosts, post)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(post)
}

func addComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Endpoint Hit: addComment")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	var body Comment
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		// return
	}
	id := vars["ID"]
	postID, _ := strconv.Atoi(id)
	for i := range TimelinePosts {
		if TimelinePosts[i].PostID == postID{
			TimelinePosts[i].Comments = append(TimelinePosts[i].Comments, body)
			json.NewEncoder(w).Encode(body)
		}
	}
}

// Favorites Requests
func returnUserFavorites(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnUserFavorites")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(UserData.Favorites)
}

func addFavorite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST Endpoint Hit: addFavorite")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody)
	
	// json.Unmarshal([]byte(reqBody), &reqBody)

	var song SongDetails
	json.Unmarshal(reqBody, &song)
	song.IsFavorite = true
	UserData.Favorites = append(UserData.Favorites, song)
	// UserFavorites = append(UserFavorites, song)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(song)
	fmt.Println(song)
}

func deleteFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	fmt.Println("DELETE Endpoint Hit: deleteFavorite")
	vars := mux.Vars(r)
	id := vars["ID"]
	favoriteID, _ := strconv.Atoi(id)
	for index, song := range UserData.Favorites {
		if song.SongID == favoriteID {
			UserData.Favorites = append(UserData.Favorites[:index], UserData.Favorites[index+1:]...)
		}
	}
}

//UserRequests
func returnUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnUser")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(UserData)
}

func main() {
	// Dummy Data for user posts
	TimelinePosts = []SongPost{
		{
			PostID: 1,
			Song : SongDetails{
				SongID : 1,
				Artist: "bty cll, Botanik",
				SongName: "Like a Drug",
				AlbumCover: "https://i.scdn.co/image/ab67616d0000b273377b5deeaf095feaa44339c1",
				IsFavorite: false,
			}, 
			Author : Author{
				Author : "Justin",
				AuthorID : 2,
			},
			Body: "Check out this song I made",
			Comments: []Comment{
				{
					CommentID: 1,
					Author: Author{
							Author: "Justin Volk",
							AuthorID: 2,
					},
					Body: "that is litty",
					PostID: 1,
				},
			},
		},
		{
			PostID: 2,
			Song : SongDetails{
				SongID : 12,
				Artist: "Louis The Child, Coin",
				SongName: "Self Care",
				AlbumCover: "https://i.scdn.co/image/ab67616d0000b2736c6c8ec19a095e0f881b9ddd",
				IsFavorite: false,
			},  
			Author : Author{
				Author : "Trevor",
				AuthorID : 1,
			},
			Body: "litty",
			Comments: []Comment{
			},
		},
		}
		UserData = UserDetails{
			UserID : 1,
			UserName : "Justin Volk",
			Favorites : []SongDetails{
				{
					SongID : 12,
					Artist: "Louis The Child, Coin",
					SongName: "Self Care",
					AlbumCover: "https://i.scdn.co/image/ab67616d0000b2736c6c8ec19a095e0f881b9ddd",
					IsFavorite: true,
				},
			},
			Following :  []string{"Trevor"},
			Followers : []string{"Trevor"},
		}
		handleRequests()
}

// Data Classes/Structures

// SongDetails is ...
type SongDetails struct {
	SongID int `json:"Song_ID"`
	Artist string `json:"Artist"`
	SongName string `json:"Song_Name"`
	AlbumCover string `json:"Album_Cover"`
	IsFavorite bool `json:"is_Favorite"`
}

// type Song struct {
// 	Song_ID int `json:"Song_ID"`
// 	Song Song_Details
// }

// Author is ...
type Author struct {
	Author string `json:"Author"`
	AuthorID int `json:"Author_ID"`
}

// Comment is ...
type Comment struct {
	CommentID int `json:"Comment_ID"`
	Author Author
	Body string `json:"Body"`
	PostID int`json:"Post_ID"`
}

// SongPost is ...
type SongPost struct {
	PostID int `json:"Post_ID"`
	Song SongDetails
	Author Author
	Body string `json:"Body"`
	Comments []Comment
}

// UserDetails is ...
type UserDetails struct {
	UserID int `json:"User_ID"`
	UserName string `json:"User_Name"`
	Favorites []SongDetails
	Following []string
	Followers []string
}

// type UserName struct {
// 	UserName string `json:UserName`
// }
// type Following struct {}
// type Followers struct {}

