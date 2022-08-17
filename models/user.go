package models

type User struct {
	ID       int      `json:"id"`
	FullName string   `json:"full_name"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     []string `json:"role"`
}
