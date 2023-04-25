package auth

import (
	"github.com/dgrijalva/jwt-go"
	"mindstore/pkg/hash-types"
)

// Define a struct to hold the JWT token
type Claims struct {
	ID *hash.Int `json:"ID"`
	jwt.StandardClaims
}
