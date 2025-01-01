package models

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name" binding:"required"`
	Role      Role   `json:"role" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=4"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
