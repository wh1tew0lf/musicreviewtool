package controllers

import (
	"encoding/json"
	"musicreviewtool/models"
	u "musicreviewtool/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var GetArtists = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllArtists(10, 0)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var CreateArtist = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	model := &models.Artist{}

	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	model.UserId = user
	resp := model.Create()
	u.Respond(w, resp)
}

var GetOneArtist = func(w http.ResponseWriter, r *http.Request) {
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

var UpdateArtist = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, convErr := strconv.Atoi(params["artistId"])
	if convErr != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	model := models.GetArtist(uint(id))

	decodeErr := json.NewDecoder(r.Body).Decode(model)
	if decodeErr != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := model.Update()
	u.Respond(w, resp)
}

var DeleteArtist = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, convErr := strconv.Atoi(params["artistId"])
	if convErr != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	model := models.GetArtist(uint(id))

	decodeErr := json.NewDecoder(r.Body).Decode(model)
	if decodeErr != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := model.Delete()
	u.Respond(w, resp)
}
