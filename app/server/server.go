package server

import (
	"github.com/urfave/negroni"
	"../routers"
)

func NewServer() *negroni.Negroni {

	server := negroni.New()
	server.UseHandler(routers.NewRouter())

	return server
}