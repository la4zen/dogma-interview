package models

type User struct {
	ID    uint   `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
}
