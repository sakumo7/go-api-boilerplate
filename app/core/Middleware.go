package core

import (
	"net/http"
)

type MiddlewareFunc func(h http.HandlerFunc) http.HandlerFunc

type Middleware func(h http.HandlerFunc) http.HandlerFunc

func Use(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}