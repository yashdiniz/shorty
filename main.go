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

	"github.com/jademaveric/shorty/internal/link"
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

	log.Println("Connecting to database...")
	db, err := sql.Open("sqlite3", "file:links.db")
	if err != nil {
		panic(err)
	}

	svc := link.NewSqliteLinkService(db, &link.DefaultHasher{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{hash}", func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r, "hash")
		link, err := svc.FindLink(hash)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// w.WriteHeader(307)
		http.Redirect(w, r, link.Target, http.StatusTemporaryRedirect)
	})

	r.Get("/add-new", func(w http.ResponseWriter, r *http.Request) {
		target := r.URL.Query().Get("target")
		if target == "" {
			w.WriteHeader(400)
			fmt.Fprintln(w, "Invalid target")
		}

		link, err := svc.AddLink(target)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, *link)
	})

	log.Println("Server started on :7000")
	log.Fatal(http.ListenAndServe(":7000", r))
}
