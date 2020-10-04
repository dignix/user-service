package domain

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	ID            uint64
	FirstName     string
	MiddleName    string
	LastName      string
	Email         string
	Password      string
	Avatar        string
	RememberToken string
	StatusID      uint32
	VerifiedAt    *timestamp.Timestamp
	CreatedAt     *timestamp.Timestamp
	UpdatedAt     *timestamp.Timestamp
}
