package link

import "fmt"

type MemoryLinkService struct {
	db     map[string](*Link)
	hasher Hasher
}

func NewMemoryLinkService(db map[string](*Link), hasher Hasher) LinkService {
	return &MemoryLinkService{db, hasher}
}

func (svc *MemoryLinkService) AddLink(link string) (*Link, error) {
	hash := svc.hasher.GenerateHash(link)
	svc.db[hash] = &Link{link, hash}
	return svc.db[hash], nil
}

func (svc *MemoryLinkService) FindLink(hash string) (*Link, error) {
	if link, ok := svc.db[hash]; ok {
		return link, nil
	} else {
		return nil, fmt.Errorf("Link with hash=%s not found", hash)
	}
}

func (svc *MemoryLinkService) DeleteLink(hash string) (*Link, error) {
	panic("MemoryLinkService::DeleteLink - Not Implemented")
}

func (svc *MemoryLinkService) ListLinks() ([](*Link), error) {
	panic("MemoryLinkService::ListLinks - Not Implemented")
}
