package controllers

import (
	"encoding/json"
	"github.com/MrTBit/goForumEngine/models"
	u "github.com/MrTBit/goForumEngine/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateThread = func(w http.ResponseWriter, r *http.Request) {
	thread := &models.Thread{}

	err := json.NewDecoder(r.Body).Decode(thread)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := thread.Create()
	u.Respond(w, resp)
}

var GetThreadsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["boardID"])
	if err != nil {
		//passed param is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetThreads(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetThread = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["threadID"])
	if err != nil {
		//passed param not int
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetThread(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)

}
