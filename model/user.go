package model

import (
	"fmt"
)

type UserRole int
type UserAccess int

const (
	// UserRole
	Invoice UserAccess = iota
	Proposal
	// UserAccess
	NoAccess UserRole = iota
	Read
	Write
)

type User struct {
	Id       int        `json:"id"`
	Fullname string     `json:"fullname"`
	Email    string     `json:"email"`
	Role     UserRole   `json"role"`
	Access   UserAccess `json:"access"`
}

func (u User) String() string {
	return fmt.Sprintf("%s partOf %s", u.Email, u.Fullname)
}
