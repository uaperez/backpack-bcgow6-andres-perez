package domain

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
} 
