package main

import (
	"log"
	"net/http"
	"shorty/api"
	_ "shorty/docs"

	"github.com/swaggo/http-swagger"
)

// @title  Shorty - URL Shortner API
// @version  0.1
// @description  This is a dummy project for the shorty url shortner
// @contact.name JadeMaveric
// @contact.url  github.com/JadeMaveric
// @host     localhost:6010
func main() {
	link_api := api.NewApi()

	http.HandleFunc("GET /{id}", link_api.RedirectLink)
	http.Handle("/link/", http.StripPrefix("/link", link_api.GetRouter()))
	http.HandleFunc("/docs/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:6010/docs/doc.json")))

	log.Println("Server listening on port :6010")
	log.Fatal(http.ListenAndServe(":6010", nil))
}
