package models

import "github.com/dgrijalva/jwt-go"
// Claim token de user
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
