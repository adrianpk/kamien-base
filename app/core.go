package app

import (
	"net/http"

	"github.com/adrianpk/kamien"
)

var (
	log = kamien.Log
)

// MethodOverride - Todo: Add comment.
func MethodOverride(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		method := r.PostFormValue("_method")
		if method == "" {
			method = r.Header.Get("X-HTTP-Method-Override")
		}
		if method == "PUT" || method == "PATCH" || method == "DELETE" {
			r.Method = method
		}
	}
}
