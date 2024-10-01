package link

import (
	"database/sql"
	"fmt"
	"log"
)

type SqliteLinkService struct {
	db     *sql.DB
	hasher Hasher
}

func NewSqliteLinkService(db *sql.DB, hasher Hasher) LinkService {
	log.Println("Setting up SqliteLinkService...")
	return &SqliteLinkService{db, hasher}
}

func (svc *SqliteLinkService) AddLink(target string) (*Link, error) {
	hash := svc.hasher.GenerateHash(target)
	link := &Link{target, hash}

	log.Printf("Adding new link %v\n", link)

	_, err := svc.db.Exec("INSERT INTO link (hash, target) VALUES (?, ?)", link.Hash, link.Target)
	if err != nil {
		return nil, fmt.Errorf("AddLink: %v", err)
	}

	return link, nil
}

func (svc *SqliteLinkService) FindLink(hash string) (*Link, error) {
	var link Link

	row := svc.db.QueryRow("SELECT hash, target FROM link WHERE hash = ?", hash)
	if err := row.Scan(&link.Hash, &link.Target); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("FindLink %s: No such link", hash)
		}
	}

	return &link, nil
}

func (svc *SqliteLinkService) DeleteLink(hash string) (*Link, error) {
	var link Link

	row := svc.db.QueryRow("DELETE FROM link WHERE hash = ? RETURNING hash, target", hash)
	if err := row.Scan(&link.Hash, &link.Target); err != nil {
		return nil, fmt.Errorf("AddLink: %v", err)
	}

	return &link, nil
}

func (svc *SqliteLinkService) ListLinks() ([](*Link), error) {
	var links [](*Link)

	rows, err := svc.db.Query("SELECT hash, target FROM link")
	if err != nil {
		return nil, fmt.Errorf("ListLinks: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var link Link
		if err := rows.Scan(&link.Hash, &link.Target); err != nil {
			return nil, fmt.Errorf("ListLinks: %v", err)
		}
		links = append(links, &link)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ListLinks: %v", err)
	}

	return links, nil
}
