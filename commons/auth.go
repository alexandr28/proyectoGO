package commons

import (
	"../models"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

var (
	privateKey *rsa.PrivateKey
	// publickey se usa paar validar el token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("no se pudo leer el archivo privado")
	}
	publickBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("no se pudo leer el archivo publico")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("no se pudo hacer el parse a privatekey")
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publickBytes)
	if err != nil {
		log.Fatal("no se pudo hacer el parse a publickey")
	}
}

// Generando  JWT
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Hour*2).Unix(),
			Issuer: "Blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("no se pudo firmar el token")
	}
	return result
}
