package main

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"net/http"
)

func main() {
	config.LoadConfig()

	mux := http.NewServeMux()
	auth.RegisterRoutes(mux)
	http.ListenAndServe(":49998", mux)
}