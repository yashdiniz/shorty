package link

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jademaveric/shorty/internal/svc/link"
)

type LinkRouter struct {
	r   *chi.Mux
	svc link.LinkService
}

func New(r *chi.Mux, svc link.LinkService) *LinkRouter {
	return &LinkRouter{r, svc}
}

func (lr *LinkRouter) GetRouter() *chi.Mux {
	r := lr.r

	r.Get("/{hash}/redirect", lr.RedirectHandler)

	// TODO: Restful routes
	r.Post("/", lr.addLinkRpcHandler)
	r.Delete("/{hash}", lr.deleteLinkRpcHandler)

	// RPC routes so I can test from the browser
	r.Get("/rpc/info", lr.getLinkRpcHandler)
	r.Get("/rpc/add", lr.addLinkRpcHandler)
	r.Get("/rpc/delete", lr.deleteLinkRpcHandler)
	r.Get("/rpc/list", lr.listLinksRpcHandler)

	return r
}

func (lr *LinkRouter) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")
	link, err := lr.svc.FindLink(hash)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, link.Target, http.StatusTemporaryRedirect) // StatusTemporaryRedirect
}

func (lr *LinkRouter) getLinkRpcHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	link, err := lr.svc.FindLink(hash)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	render.Render(w, r, NewLinkResponse(link))
}

func (lr *LinkRouter) addLinkRpcHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Invalid target")
	}

	link, err := lr.svc.AddLink(target)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	render.Render(w, r, NewLinkResponse(link))
}

func (lr *LinkRouter) deleteLinkRpcHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get("hash")
	if hash == "" {
		http.Error(w, "Missing hash", http.StatusBadRequest)
	}

	link, err := lr.svc.DeleteLink(hash)
	if err != nil {
		http.Error(w, fmt.Sprintf("Hash %s not found", hash), http.StatusNotFound)
		return
	}

	render.Render(w, r, NewLinkResponse(link))
}

func (lr *LinkRouter) listLinksRpcHandler(w http.ResponseWriter, r *http.Request) {
	links, err := lr.svc.ListLinks()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch links: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	resp := make([]render.Renderer, len(links))
	for i, l := range links {
		resp[i] = NewLinkResponse(l)
	}

	render.RenderList(w, r, resp)
}
