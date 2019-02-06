package controllers

import (
	"../commons"
	"../models"
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
)

//ValidateToken permite validar el token del cliente

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m models.Message
	/*
		token, err := request.ParseFromRequestWithClaims(r,request.OAuth2Extractor,&models.Claim{}, func(token *jwt.Token) ( interface{},  error) {
			return commons.PublicKey,nil
		})*/

	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return commons.PublicKey, nil
	}, request.WithClaims(&models.Claim{}))
	if err != nil {
		m.CodState = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			valError := err.(*jwt.ValidationError)
			switch valError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "su token a expirado"
				commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "la firma del token no coincide"
				commons.DisplayMessage(w, m)
				return
			default:
				m.Message = "su token no es valido"
				commons.DisplayMessage(w, m)
				return
			}
		}
	}
	if token.Valid {
		// obtenemos en el contexto token del usuario
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		next(w, r.WithContext(ctx))
	} else {
		m.CodState = http.StatusUnauthorized
		m.Message = "su token no es valido"
		commons.DisplayMessage(w, m)
	}
}
