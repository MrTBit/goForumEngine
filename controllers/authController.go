package controllers

import (
	"encoding/json"
	"github.com/MrTBit/goForumEngine/models"
	u "github.com/MrTBit/goForumEngine/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode body into struct and fail if any errors occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))
		return
	}

	resp := account.Create() //create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any errors occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
