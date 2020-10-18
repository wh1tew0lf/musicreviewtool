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

	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/artists", controllers.CreateArtist).Methods("POST")
	router.HandleFunc("/api/artists", controllers.GetArtists).Methods("GET")
	router.HandleFunc("/api/artists/{artistId}", controllers.GetOneArtist).Methods("GET")
	router.HandleFunc("/api/artists/{artistId}", controllers.UpdateArtist).Methods("PUT")
	router.HandleFunc("/api/artists/{artistId}", controllers.DeleteArtist).Methods("DELETE")

	router.HandleFunc("/api/artists/{artistId}/albums", controllers.CreateAlbum).Methods("POST")
	router.HandleFunc("/api/artists/{artistId}/albums/{albumId}", controllers.GetOneAlbum).Methods("GET")
	router.HandleFunc("/api/artists/{artistId}/albums/{albumId}", controllers.UpdateAlbum).Methods("PUT")
	router.HandleFunc("/api/artists/{artistId}/albums/{albumId}", controllers.DeleteAlbum).Methods("DELETE")

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
