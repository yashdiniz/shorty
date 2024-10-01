package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jademaveric/shorty/internal/database"
	LinkRouter "github.com/jademaveric/shorty/internal/routes/link"
	"github.com/jademaveric/shorty/internal/svc/link"
)

func main() {
	log.Println("Server is starting up")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.OpenSqliteDb()
	svc := link.NewSqliteLinkService(db, &link.DefaultHasher{})
	linkRouter := LinkRouter.New(chi.NewRouter(), svc)

	r.Mount("/link", linkRouter.GetRouter())

	r.Get("/{hash}", linkRouter.RedirectHandler)

	log.Println("Server started on :7000")
	log.Fatal(http.ListenAndServe(":7000", r))
}
