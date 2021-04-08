package domain

import (
	"encoding/json"
	"time"
)

type (
	Domain struct {
		Domain    string `json:"domain"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func NewDomain(domain string) *Domain {
	now := time.Now()
	return &Domain{
		Domain:   domain,
		CreateAt: now,
	}
}

func (d *Domain) ToJson() ([]byte, error) {
	str, err := json.Marshal(d)
	if err != nil {
		return []byte("[]"), err
	}
	return str, nil
}
