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
var SearchResults []Song_Details
var TimelinePosts []SongPost
var UserFavorites []Song_Details
var UserData User_Details



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
	myRouter.HandleFunc("/search", searchItunesForArtistId)
	myRouter.HandleFunc("/User", returnUser)
	myRouter.HandleFunc("/favorites", returnUserFavorites)
	myRouter.HandleFunc("/favorites", addFavorite).Methods("POST")
	myRouter.HandleFunc("/favorites/{ID}", deleteFavorite).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

//Search Results Requests



func searchItunesForArtistId(w http.ResponseWriter, r *http.Request) {
	var url = `https://itunes.apple.com/search?term=btycll&country=US&entity=song,album,podcast`
	resp, err := http.Get(url)
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
		if song.Song_ID == (keyInt) {
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
		if post.Post_ID == (keyInt) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(post)
		}
	}
}

func createNewSongPost(w http.ResponseWriter, r *http.Request){
	fmt.Println("POST Endpoint Hit: returnTimelinePosts")
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
		if TimelinePosts[i].Post_ID == postID{
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
	reqBody, _ := ioutil.ReadAll(r.Body)
	var song Song_Details
	json.Unmarshal(reqBody, &song)
	UserFavorites = append(UserData.Favorites, song)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(song)
}

func deleteFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	favoriteID, _ := strconv.Atoi(id)
	for index, song := range UserFavorites {
		if song.Song_ID == favoriteID {
			UserFavorites = append(UserData.Favorites[:index], UserData.Favorites[index+1:]...)
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
	// dummy data for search results
	SearchResults = []Song_Details{
		Song_Details{
			Song_ID: 1,
			Artist: "bty cll, Botanik",
			Song_Name: "Like a Drug",
			Album_Cover: "https://i.scdn.co/image/ab67616d0000b273377b5deeaf095feaa44339c1",
			IsFavorite: false,
		},
		Song_Details{
			Song_ID: 5,
			Artist: "Louis The Child, Drew Love",
			Song_Name: "Free",
			Album_Cover: "https://i.scdn.co/image/ab67616d0000b273d0c97444ecc52c4ca601144a",
			IsFavorite: false,
		},
			Song_Details{
			Song_ID: 6,
			Artist: "bty cll",
			Song_Name: "Here Alone",
			Album_Cover: "https://m.media-amazon.com/images/I/71SFywf-m9L._SS500_.jpg",
			IsFavorite: false,
		},
		Song_Details{
			Song_ID: 7,
			Artist: "Elohim",
			Song_Name: "Sensations - Whethan remix",
			Album_Cover: "https://i.scdn.co/image/ab67616d0000b273b708f022a637cf80ec2f7c57",
			IsFavorite: false,
		},
	
	}
	// Dummy Data for user posts
	TimelinePosts = []SongPost{
		SongPost{
			Post_ID: 1,
			Song : Song_Details{
				Song_ID : 1,
				Artist: "bty cll, Botanik",
				Song_Name: "Like a Drug",
				Album_Cover: "https://i.scdn.co/image/ab67616d0000b273377b5deeaf095feaa44339c1",
				IsFavorite: false,
			}, 
			Author : Author{
				Author : "Justin",
				Author_ID : 2,
			},
			Body: "Check out this song I made",
			Comments: []Comment{
				Comment{
					Comment_ID: 1,
					Author: Author{
							Author: "Justin Volk",
							Author_ID: 2,
					},
					Body: "that is litty",
					Post_ID: 1,
				},
			},
		},
		SongPost{
			Post_ID: 2,
			Song : Song_Details{
				Song_ID : 12,
				Artist: "Louis The Child, Coin",
				Song_Name: "Self Care",
				Album_Cover: "https://i.scdn.co/image/ab67616d0000b2736c6c8ec19a095e0f881b9ddd",
				IsFavorite: false,
			},  
			Author : Author{
				Author : "Trevor",
				Author_ID : 1,
			},
			Body: "litty",
			Comments: []Comment{
				Comment{},
				},
			},
		}
		UserData = User_Details{
			User_ID : 1,
			User_Name : "Justin Volk",
			Favorites : []Song_Details{
				{
					Song_ID : 12,
					Artist: "Louis The Child, Coin",
					Song_Name: "Self Care",
					Album_Cover: "https://i.scdn.co/image/ab67616d0000b2736c6c8ec19a095e0f881b9ddd",
					IsFavorite: true,
				},
			},
			Following :  []string{"Trevor"},
			Followers : []string{"Trevor"},
		}
		handleRequests()
}

// Data Classes/Structures
type Song_Details struct {
	Song_ID int `json:"Song_ID"`
	Artist string `json:"Artist"`
	Song_Name string `json:"Song_Name"`
	Album_Cover string `json:"Album_Cover"`
	IsFavorite bool `json:"isFavorite"`
}

// type Song struct {
// 	Song_ID int `json:"Song_ID"`
// 	Song Song_Details
// }

type Author struct {
	Author string `json:"Author"`
	Author_ID int `json:"Author_ID"`
}

type Comment struct {
	Comment_ID int `json:"Comment_ID"`
	Author Author
	Body string `json:"Body"`
	Post_ID int`json:"Post_ID"`
}

type SongPost struct {
	Post_ID int `json:"Post_ID"`
	Song Song_Details
	Author Author
	Body string `json:"Body"`
	Comments []Comment
}

type User_Details struct {
	User_ID int `json:"User_ID"`
	User_Name string `json:"User_Name"`
	Favorites []Song_Details
	Following []string
	Followers []string
}

// type UserName struct {
// 	UserName string `json:UserName`
// }
// type Following struct {}
// type Followers struct {}

