package controllers

import (
	"encoding/json"
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

	user := r.Context().Value("user").(uint)
	model := &models.Rating{}

	decodeErr := json.NewDecoder(r.Body).Decode(model)
	if decodeErr != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := model.RateArtist(user, uint(id))
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
