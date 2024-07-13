package link

import (
	"log"
	"shorty/db"
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
	dal db.LinkDAL
}

func New() LinkService {
	dal, err := db.NewLinkDAL()
	if err != nil {
		log.Panicf("Cannot connect to database: %w\n", err)
	}

	svc := linkSvcImpl{
		dal: dal,
	}

	return &svc
}

func (s *linkSvcImpl) CreateLink(params LinkCreateParams) *Link {
  l, err := s.dal.Insert(db.LinkModel{
    Id: utils.GenerateKey(4),
    Name: params.Name,
    Target: params.Target,
    Ctime: utils.GetISOTimestamp(nil),
    Mtime: utils.GetISOTimestamp(nil),
    Visits: 0,
  })

  if err == nil {
    log.Panicln()("Could not save")
  }

	// return &l
}

func (s *linkSvcImpl) GetAllLinks() [](*Link) {
	panic("TODO")
}

func (s *linkSvcImpl) GetLink() *Link {
	panic("TODO")
}
