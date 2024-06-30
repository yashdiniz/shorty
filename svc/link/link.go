package link

import (
	"shorty/internal/utils"
)

type Link struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Target string `json:"target"`
	Ctime  string `json:"ctime"`
	Mtime  string `json:"mtime"`
	Visits int    `json:"visits"`
}

type LinkCreateParams struct {
	Name   string `json:"name"`
	Target string `json:"target"`
}

type LinkService interface {
	CreateLink(params LinkCreateParams) *Link
	GetAllLinks() [](*Link)
	GetLink() *Link
}

type linkSvcImpl struct {
	db [](*Link)
}

func New() LinkService {
	svc := linkSvcImpl{
		db: make([]*Link, 0),
	}

	return &svc
}

func (s *linkSvcImpl) CreateLink(params LinkCreateParams) *Link {
	l := Link{
		Name:   params.Name,
		Target: params.Target,
	}

	l.Id = utils.GenerateKey(4)
	l.Ctime = utils.GetISOTimestamp(nil)
	l.Mtime = utils.GetISOTimestamp(nil)

	// TODO: Use a db wrapper
	s.db = append(s.db, &l)

	return &l
}

func (s *linkSvcImpl) GetAllLinks() [](*Link) {
	panic("TODO")
}

func (s *linkSvcImpl) GetLink() *Link {
	panic("TODO")
}
