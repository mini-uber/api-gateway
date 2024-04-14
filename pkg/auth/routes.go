package auth

import "net/http"

const basePath = "api/v1/auth"

func RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc(basePath+"/login", Login)
	r.HandleFunc(basePath+"/register", Register)
}