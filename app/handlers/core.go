package handlers

import (
	"../core"
	"../models"
	"net/http"
)

func CreateGetAllHandler(m models.Model) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		res := core.Response{ResponseWriter: w}
		//obj := new(m)
		//objArray :=
		res.SendOK(m.FetchAll())
	}
}
