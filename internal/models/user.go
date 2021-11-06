package models

type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
