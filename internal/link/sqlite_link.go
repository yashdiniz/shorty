package link

import (
	"database/sql"
	"fmt"
)

type SqliteLinkService struct {
	db     *sql.DB
	hasher Hasher
}

func NewSqliteLinkService(db *sql.DB, hasher Hasher) LinkService {
	return &SqliteLinkService{db, hasher}
}

func (svc *SqliteLinkService) AddLink(target string) (*Link, error) {
	hash := svc.hasher.GenerateHash(target)
	link := &Link{target, hash}

	_, err := svc.db.Exec("INSERT INTO link (hash, target) VALUES (?, ?)", link.hash, link.target)
	if err != nil {
		return nil, fmt.Errorf("AddLink: %v", err)
	}

	return link, nil
}

func (svc *SqliteLinkService) FindLink(hash string) (*Link, error) {
	var link Link

	row := svc.db.QueryRow("SELECT hash, target FROM link WHERE hash = ?", hash)
	if err := row.Scan(&link.hash, &link.target); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("FindLink %s: No such link", hash)
		}
	}

	return &link, nil
}
