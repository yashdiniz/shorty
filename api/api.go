package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"shorty/svc/link"
)

type ErrorRes struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

type Api interface {
	GetRouter() *http.ServeMux
	RedirectLink(w http.ResponseWriter, r *http.Request)
}

type listApi struct {
	db     [](*link.Link)
	svc    link.LinkService
	logger *log.Logger
}

func NewApi() Api {
	api := &listApi{
		db:     make([](*link.Link), 0),
		svc:    link.New(),
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	// Sample data for testing
	dev_links := [](*link.Link){
		&link.Link{Id: "abc", Name: "test1", Target: "https://google.com"},
		&link.Link{Id: "123", Name: "test2", Target: "https://google.com"},
		&link.Link{Id: "zyx", Name: "test3", Target: "https://google.com"},
	}
	api.db = append(api.db, dev_links...)

	return api
}

func (api *listApi) GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", api.getAllLinks)
	mux.HandleFunc("GET /{id}", api.getLink)
	mux.HandleFunc("POST /{$}", api.createLink)
	return mux
}

func (api *listApi) createLink(w http.ResponseWriter, r *http.Request) {
	api.logger.Printf("Creating new link\n")

	var params link.LinkCreateParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newLink := api.svc.CreateLink(params)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newLink)
}

func (api *listApi) getAllLinks(w http.ResponseWriter, r *http.Request) {
	api.logger.Printf("Fetching all links\n")
	res := api.db
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	return
}

func (api *listApi) getLink(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	api.logger.Printf("Fetching link %s\n", id)

	// search db for id
	for _, l := range api.db {
		if l.Id == id {
			res := l
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	res := ErrorRes{"Not Found", fmt.Sprintf("Link %s NOT found", id)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}

func (api *listApi) RedirectLink(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	api.logger.Printf("Redirecting to link %s\n", id)

	// search db for id
	for _, l := range api.db {
		if l.Id == id {
			l.Visits++
			l.Mtime = fmt.Sprint(time.Now().UTC().Format(time.RFC822))
			http.Redirect(w, r, l.Target, http.StatusSeeOther)
			return
		}
	}

	res := ErrorRes{"Not Found", fmt.Sprintf("Link %s NOT found", id)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
