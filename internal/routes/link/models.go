package link

import (
	"fmt"
	"net/http"

	"github.com/jademaveric/shorty/internal/svc/link"
)

type LinkResponse struct {
	Target string `json:"target"`
	Hash   string `json:"hash"`
	Href   string `json:"href"`
}

func NewLinkResponse(link *link.Link) *LinkResponse {
	return &LinkResponse{link.Target, link.Hash, ""}
}

func (l *LinkResponse) Render(w http.ResponseWriter, r *http.Request) error {
	l.Href = fmt.Sprintf("http://%s/link/rpc/info?hash=%s", r.Host, l.Hash)
	return nil
}
