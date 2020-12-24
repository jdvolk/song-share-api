## About
  - This is an api for a personal site that will allow you to search for a song and create a post about it
  - It is written in goLang
  - it runs on localhost:10000
## End Points
- "/" 
  -  welcome page
- "/Home"
  - Returns a timeline of dummy posts
```
  [
    {
      "Post_ID": 1,
      "Song": {
      "Song_ID": 1,
      "Song": {
        "Artist": "bty cll, Botanik",
        "Song_Name": "Like a Drug",
        "Album_Cover": "https://i.scdn.co/image/ab67616d0000b273377b5deeaf095feaa44339c1",
        isFavorite: false,
        },
        
      },
      "Author": {
        "Author": "Justin",
        "Author_ID": 2
      },
      "Body": "Check out this song I made",
      "Comments": null
    },
    {
      "Post_ID": 2,
      "Song": {
      "Song_ID": 5,
      "Song": {
        "Artist": "Louis The Child, Coin",
        "Song_Name": "Self Care",
        "Album_Cover": "https://i.scdn.co/image/ab67616d0000b2736c6c8ec19a095e0f881b9ddd",
        isFavorite: false,
        }
      },
      "Author": {
        "Author": "Trevor",
        "Author_ID": 1
      },
      "Body": "Check out this song I made",
      "Comments": null
    }
  ]
```
  
-  "/Home/{ID}"
   -  returns a single post
  ```
  {
    "Post_ID": 2,
    "Song": {
    "Song_ID": 5,
    "Song": {
      "Artist": "Louis The Child, Coin",
      "Song_Name": "Self Care",
      "Album_Cover": "https://i.scdn.co/image/ab67616d0000b2736c6c8ec19a095e0f881b9ddd",
      isFavorite: false,
      }
    },
    "Author": {
      "Author": "Trevor",
      "Author_ID": 1
    },
    "Body": "Check out this song I made",
    "Comments": null
  }
  ```
   -  PUT to add a comment to a post
-  "/searchResults" 
   -  returns a dummy list of search results
```
  [
    {
    "Song_ID": 1,
    "Song": {
      "Artist": "bty cll, Botanik",
      "Song_Name": "Like a Drug",
      "Album_Cover": "https://i.scdn.co/image/ab67616d0000b273377b5deeaf095feaa44339c1",
      isFavorite: false,
      }
    },
    {
    "Song_ID": 5,
    "Song": {
      "Artist": "Louis The Child, Drew Love",
      "Song_Name": "Free",
      "Album_Cover": "https://i.scdn.co/image/ab67616d0000b273d0c97444ecc52c4ca601144a",
      isFavorite: false,
      }
    },
    {
    "Song_ID": 6,
    "Song": {
      "Artist": "bty cll",
      "Song_Name": "Here Alone",
      "Album_Cover": "https://m.media-amazon.com/images/I/71SFywf-m9L._SS500_.jpg",
      isFavorite: false,
      }
    },
    {
    "Song_ID": 7,
    "Song": {
      "Artist": "Elohim",
      "Song_Name": "Sensations - Whethan remix",
      "Album_Cover": "https://i.scdn.co/image/ab67616d0000b273b708f022a637cf80ec2f7c57",
      isFavorite: false,
      }
    }
  ]
```
-  "/searchResults/{ID}"
   -  Returns a single search result
```
  {
  "Song_ID": 7,
  "Song": {
    "Artist": "Elohim",
    "Song_Name": "Sensations - Whethan remix",
    "Album_Cover": "https://i.scdn.co/image/ab67616d0000b273b708f022a637cf80ec2f7c57",
    isFavorite: false,
    }
  }
```
   -  POST to this endpoint will create a user post about this result
-  "favorites"
   -  returns all userFavorites
   -  POST will add a favorite to the users list
-  "/favorites/{ID}"
   -  Delete a favorite from the users list


