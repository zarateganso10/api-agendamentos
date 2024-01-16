package entity

type Auth struct {
	Email    string `json:"email"`
	Type     string `json:"type"`
	Password string `json:"password"`
}
