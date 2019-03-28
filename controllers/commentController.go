package controllers

import (
	"encoding/json"
	"github.com/MrTBit/goForumEngine/models"
	u "github.com/MrTBit/goForumEngine/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var CreateComment = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["threadID"])
	if err != nil {
		//passed param is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	comment := &models.Comment{}
	comment.ThreadID = uint(id)

	err = json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := comment.Create()
	u.Respond(w, resp)
}

var GetCommentsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["threadID"])
	if err != nil {
		//passed param is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetComments(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetComment = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["postID"])
	if err != nil {
		//passed param not int
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetComment(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)

}
