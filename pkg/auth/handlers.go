package auth

import (
	"api-gateway/pkg/config"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post(
		config.LoadedConfig.AUTH_ADDR+"/api/v1/auth/login", 
		"application/json", 
		r.Body,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func Register(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post(
		config.LoadedConfig.AUTH_ADDR+"/api/v1/auth/register", 
		"application/json", 
		r.Body,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
