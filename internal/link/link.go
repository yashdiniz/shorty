package link

import (
	"encoding/base64"
	"encoding/binary"
	"hash/fnv"
)

type LinkService interface {
	AddLink(link string) (*Link, error)
	FindLink(hash string) (*Link, error)
}

type Link struct {
	target string
	hash string
}

type Hasher interface {
	GenerateHash(link string) string
}

type DefaultHasher struct{}

func (h *DefaultHasher) GenerateHash(link string) string {
	hash := fnv.New32()
	hash.Write([]byte(link))
	val := hash.Sum32()

	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, val)

	return base64.RawURLEncoding.EncodeToString(buffer)
}