package models

import ()

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}