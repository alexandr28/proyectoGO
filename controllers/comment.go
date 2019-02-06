package controllers

import (
	"../commons"
	"../configuration"
	"../models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CommentCreate(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	user:= models.User{}
	m := models.Message{}

	user, _= r.Context().Value("user").(models.User)

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.CodState = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	comment.UserID=user.ID

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.CodState = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}
	m.CodState = http.StatusCreated
	m.Message = "Comentario creado con exito"
	commons.DisplayMessage(w, m)
}

// CommentGetAll obtiene todos los comentarios
func CommentGetAll(w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	vote := models.Vote{}
	user, _ =r.Context().Value("user").(models.User)
	vars := r.URL.Query()
	db := configuration.GetConnection()
	defer db.Close()
	consultComents := db.Where("parent_id=0")
	if order, ok := vars["order"]; ok {
		if order[0] == "votes" {
			consultComents = consultComents.Order("votes desc, created_at desc")
		}
	} else {
		if idlimit, ok := vars["idlimit"]; ok {
			registterByPage := 30
			offset, err := strconv.Atoi(idlimit[0])
			if err != nil {
				log.Println("error")
			}
			consultComents = consultComents.Where("id BETWEEN ? AND", offset-registterByPage)
		}
		consultComents = consultComents.Order("id desc")
	}
	consultComents.Find(&comments)
	for i := range comments {
		db.Model(&comments[i]).Related(&comments[i].User)
		comments[i].User[0].Password = ""
		comments[i].Children = commentGetChildren(comments[i].ID)

		// se busca el voto del usuario que esta en sesion
		vote.CommentID = comments[i].ID
		vote.UserID = user.ID
		count := db.Where(&vote).Find(&vote).RowsAffected
		if count > 0 {
			if vote.Value {
				comments[i].HasVote = 1
			} else {
				comments[i].HasVote = -1
			}
		}

	}
	j, err := json.Marshal(comments)
	if err != nil {
		m.CodState = http.StatusInternalServerError
		m.Message = "Error al convertir los comentarios en json"
		commons.DisplayMessage(w, m)
		return
	}
	if len(comments) > 0 {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m.CodState = http.StatusNoContent
		m.Message = "no se encontraron comentarios"
		commons.DisplayMessage(w, m)

	}
}

func commentGetChildren(id uint) (children []models.Comment) {
	db := configuration.GetConnection()
	defer db.Close()
	db.Where("parent_id=?", id).Find(&children)
	for i := range children {
		db.Model(&children[i]).Related(&children[i].User)
		children[i].User[0].Password = ""
	}
	return
}
