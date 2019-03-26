package controllers

import (
	"encoding/json"
	"github.com/MrTBit/goForumEngine/models"
	u "github.com/MrTBit/goForumEngine/utils"
	"net/http"
)

var CreateBoard = func(w http.ResponseWriter, r *http.Request) {
	board := &models.Board{}

	err := json.NewDecoder(r.Body).Decode(board)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := board.Create()
	u.Respond(w, resp)
}

var GetBoards = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetBoards()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
