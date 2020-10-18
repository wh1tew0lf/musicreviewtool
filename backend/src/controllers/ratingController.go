package controllers

import (
	"musicreviewtool/models"
	u "musicreviewtool/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var RateArtist = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["artistId"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetArtist(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var RateAlbum = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["albumId"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetArtist(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
