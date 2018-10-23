package handlers

import (
	"net/http"
	"os"
	"../core"
)

type APIInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func GetAPIInfo(w http.ResponseWriter, r *http.Request) {
	res := core.Response{ResponseWriter: w}
	res.SendOK(APIInfo{
		Name:    os.Getenv("API_TITLE"),
		Version: os.Getenv("API_VERSION"),
	})
}