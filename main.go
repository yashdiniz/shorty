package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

  _ "github.com/mattn/go-sqlite3"
	"github.com/jademaveric/shorty/internal/link"
)

// TEST
func main() {
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
