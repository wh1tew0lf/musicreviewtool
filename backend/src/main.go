package main

import (
	"fmt"
	"musicreviewtool/app"
	"musicreviewtool/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/artists", controllers.CreateArtist).Methods("POST")
	router.HandleFunc("/artists", controllers.GetArtists).Methods("GET")
	router.HandleFunc("/artists/{artistId}", controllers.GetOneArtist).Methods("GET")
	router.HandleFunc("/artists/{artistId}", controllers.UpdateArtist).Methods("PUT")
	router.HandleFunc("/artists/{artistId}", controllers.DeleteArtist).Methods("DELETE")

	router.HandleFunc("/artists/{artistId}/albums", controllers.CreateAlbum).Methods("POST")
	router.HandleFunc("/artists/{artistId}/albums/{albumId}", controllers.GetOneAlbum).Methods("GET")
	router.HandleFunc("/artists/{artistId}/albums/{albumId}", controllers.UpdateAlbum).Methods("PUT")
	router.HandleFunc("/artists/{artistId}/albums/{albumId}", controllers.DeleteAlbum).Methods("DELETE")

	router.HandleFunc("/artists/{artistId}/rate", controllers.RateArtist).Methods("PUT")
	router.HandleFunc("/artists/{artistId}/albums/{albumId}/rate", controllers.RateAlbum).Methods("PUT")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
