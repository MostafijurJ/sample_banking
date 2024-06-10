package serializer

import (
	"github.com/pelletier/go-toml/v2"
)

type User struct {
	Username string         `json:"username"`
	FullName string         `json:"full_name"`
	Email    string         `json:"email"`
	Dob      toml.LocalDate `json:"dob"`
}
