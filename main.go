package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jademaveric/shorty/internal/database"
	LinkRouter "github.com/jademaveric/shorty/internal/routes/link"
	"github.com/jademaveric/shorty/internal/svc/link"
)

// TEST
func mainOG() {
	// svc := link.NewMemoryLinkService(make(map[string]*link.Link), &link.DefaultHasher{})

	db, err := sql.Open("sqlite3", "file:links.db")
	if err != nil {
		panic(err)
	}
	svc := link.NewSqliteLinkService(db, &link.DefaultHasher{})

	for true {
		fmt.Print("Choose:\n1. add <url>\n2. find <hash>\n3. exit\n> ")
		var (
			inp string
			arg string
		)
		fmt.Scan(&inp)

		if strings.HasPrefix(inp, "add") {
			fmt.Scan(&arg)
			link, _ := svc.AddLink(arg)
			fmt.Println(link)
		} else if strings.HasPrefix(inp, "find") {
			fmt.Scan(&arg)
			link, _ := svc.FindLink(arg)
			fmt.Println(link)
		} else if strings.HasPrefix(inp, "exit") {
			os.Exit(0)
		} else {
			fmt.Println("Unknown option")
		}
	}
}

func main() {
	log.Println("Server is starting up")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.OpenSqliteDb()
	svc := link.NewSqliteLinkService(db, &link.DefaultHasher{})
	lr := LinkRouter.New(chi.NewRouter(), svc)

	r.Mount("/link", lr.GetRouter())
	r.Get("/{hash}", lr.RedirectHandler)

	log.Println("Server started on :7000")
	log.Fatal(http.ListenAndServe(":7000", r))
}
