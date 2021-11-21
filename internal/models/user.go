package models

type User struct {
	Id       int    `json:"id,omitempty"`
	AuthType string `json:"auth_type"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
