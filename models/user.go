package models

import (
	"encoding/gob"
)

func init() {
	gob.Register(&User{})
}

// User - represents a user in Meshery
type User struct {
	UserID    string `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}
