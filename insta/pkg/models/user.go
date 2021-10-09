package models

type User struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}
