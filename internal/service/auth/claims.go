package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// Define a struct to hold the JWT token
type Claims struct {
	ID int `json:"ID"`
	jwt.StandardClaims
}
