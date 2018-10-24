package handlers

import (
	"../core"
	"../models"
	"net/http"
)

func GetAllUsersHandler(w http.ResponseWriter, req *http.Request) {
	res := core.Response{ResponseWriter: w}
	user := new(models.User)
	users := user.FetchAll()
	res.SendOK(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	user := new(models.User)
	req.GetJSONBody(user)

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendCreated(user)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := models.User{
		ID: id,
	}

	if err := user.FetchById(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendOK(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	id, _ := req.GetVarID()

	user := new(models.User)
	req.GetJSONBody(user)
	user.ID = id

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendOK(user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	req := core.Request{ResponseWriter: w, Request: r}
	res := core.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := models.User{
		ID: id,
	}

	if err := user.Delete(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendNoContent()
}
