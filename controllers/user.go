package controllers

import (
	"../commons"
	"../configuration"
	"../models"
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	db := configuration.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)

	db.Where("email =? and password=?", user.Email, pwd).First(&user)
	if user.ID > 0 {
		user.Password = ""
		token := commons.GenerateJWT(user)
		j, err := json.Marshal(models.Token{Token: token})
		if err != nil {
			log.Fatalf("Error al convertir a JSON:  %s", err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m := models.Message{
			Message:  "Usuario o Password no valido",
			CodState: http.StatusUnauthorized,
		}
		commons.DisplayMessage(w, m)
	}
}

// UserCreate Permit register un user
func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	m := models.Message{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Sprintf("Error al leer el usuario a registrar: %s", err)
		m.CodState = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	if user.Password != user.ConfirmPassword {
		m.Message = "Las contraseñas no coinciden"
		m.CodState = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	c := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", c)
	user.Password = pwd
	// usamos el email para codificar nuestra imagen en md5
	picturemd5 := md5.Sum([]byte(user.Email))
	// string codificado en md5
	picstr := fmt.Sprintf("%x", picturemd5)
	// pasamos a gravatar nuesto string codificado
	user.Picture = "https://gravatar.com/avatar/" + picstr + "?s=100"

	db := configuration.GetConnection()
	defer db.Close()
	err = db.Create(&user).Error
	if err != nil {
		m.Message = fmt.Sprintf("error al crear el registro %s", err)
		m.CodState = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}

	m.Message = "Usuario creado con exito "
	m.CodState = http.StatusCreated
	commons.DisplayMessage(w, m)
}
