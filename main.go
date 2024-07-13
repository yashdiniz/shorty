package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
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
	router := httprouter.New()

	router.GET("/", Index)
  router.GET("/hello/:name", Hello)

  log.Fatal(http.ListenAndServe(":7000", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
