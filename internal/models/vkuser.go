package models

type VKUser struct {
	VKId     int    `json:"vkid"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email,omitempty"`
}
