package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var songs []Song = make([]Song, 0, 20)

type Song struct {
	ID     string  `json:"id"`
	Sid    string  `json:"sid"`
	Title  string  `json:"title"`
	Artist *Artist `json:"artist"`
}

type Artist struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	// Just to print it
	// for _, i := range songs {
	// 	fmt.Fprintf(w, "id:%v\nsid:%v\ntitle:%v\nartist:%v %v\n", i.ID, i.Sid, i.Title, i.Artist.Firstname, i.Artist.Lastname)
	// 	fmt.Fprintln(w, "---------------------------------------")
	// }

	w.Header().Set("Content-Type", "applications/json")
	json.NewEncoder(w).Encode(songs)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r)

	for index, item := range songs {
		if item.ID == params["id"] {
			songs = append(songs[:index], songs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(songs)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	params := mux.Vars(r)

	for _, i := range songs {
		if i.ID == params["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	song := Song{}
	_ = json.NewDecoder(r.Body).Decode(&song)
	song.ID = strconv.Itoa(rand.Intn(100))
	songs = append(songs, song)

	json.NewEncoder(w).Encode(songs)
}

func updateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// deleted value, but lets try to change the value of the struct directly
	// for index, item := range songs {
	// 	if params["id"] == item.ID {
	// 		songs = append(songs[:index], songs[index+1:]...)
	// 		break
	// 	}
	// }

	song := Song{}
	_ = json.NewDecoder(r.Body).Decode(&song)

	for index, item := range songs {
		if params["id"] == item.ID {
			songs[index].Sid = song.Sid
			songs[index].Title = song.Title
			songs[index].Artist.Firstname = song.Artist.Firstname
			songs[index].Artist.Lastname = song.Artist.Lastname
			break
		}
	}

	json.NewEncoder(w).Encode(songs)
}

func main() {
	const port = ":8000"
	// Too lazy, justin isnt my favourite artist, gpt gave me these songs
	songs = append(songs, Song{"1", "2000", "Baby", &Artist{Firstname: "Justin", Lastname: "Beiber"}})
	songs = append(songs, Song{"2", "2001", "Love Yourself", &Artist{Firstname: "Justin", Lastname: "Beiber"}})
	songs = append(songs, Song{"3", "2002", "Sorry", &Artist{Firstname: "Justin", Lastname: "Beiber"}})

	fmt.Printf("%v", songs[0].Artist.Firstname)
	r := mux.NewRouter()

	r.HandleFunc("/songs", getSongs).Methods("GET")
	r.HandleFunc("/songs/{id}", getSong).Methods("GET")
	r.HandleFunc("/songs", createSong).Methods("POST")
	r.HandleFunc("/songs/{id}", updateSong).Methods("PUT")
	r.HandleFunc("/songs/{id}", deleteSong).Methods("DELETE")

	log.Fatal(http.ListenAndServe(port, r))

}
