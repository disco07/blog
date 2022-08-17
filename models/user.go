package models

type User struct {
	ID       int      `json:"id" db:"id"`
	FullName string   `json:"full_name" db:"full_name"`
	Username string   `json:"username" db:"username"`
	Email    string   `json:"email" db:"email"`
	Password string   `json:"password" db:"password"`
	Role     []string `json:"role" db:"role"`
}
