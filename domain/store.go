package domain

type Store struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsOpen      bool   `json:"is_open"`
}
