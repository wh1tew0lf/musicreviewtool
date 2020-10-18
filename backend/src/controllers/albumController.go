package controllers

import (
	"encoding/json"
	"musicreviewtool/models"
	u "musicreviewtool/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateAlbum = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, convertErr := strconv.Atoi(params["artistId"])
	if convertErr != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	user := r.Context().Value("user").(uint)
	model := &models.Album{}

	decodeErr := json.NewDecoder(r.Body).Decode(model)
	if decodeErr != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	model.ArtistId = uint(id)
	model.UserId = user
	resp := model.Create()
	u.Respond(w, resp)
}

var GetOneAlbum = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["albumId"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetAlbum(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateAlbum = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, convErr := strconv.Atoi(params["albumId"])
	if convErr != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	model := models.GetAlbum(uint(id))

	decodeErr := json.NewDecoder(r.Body).Decode(model)
	if decodeErr != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := model.Update()
	u.Respond(w, resp)
}

var DeleteAlbum = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, convErr := strconv.Atoi(params["albumId"])
	if convErr != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	model := models.GetAlbum(uint(id))

	decodeErr := json.NewDecoder(r.Body).Decode(model)
	if decodeErr != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := model.Delete()
	u.Respond(w, resp)
}
