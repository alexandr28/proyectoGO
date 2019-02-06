package controllers

import (
	"../commons"
	"../configuration"
	"../models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el usuario a registrar: %s", err)
		m.CodState = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	vote.UserID = user.ID
	db := configuration.GetConnection()
	defer db.Close()
	db.Where("comment_id=? and user_id=?", vote.CommentID, vote.UserID).First(&currentVote)
	// si no existe
	if currentVote.ID == 0 {
		db.Create(&vote)
		err := updateCommentVotes(vote.CommentID, vote.Value, false)
		if err != nil {
			m.Message = err.Error()
			m.CodState = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "voto registrado"
		m.CodState = http.StatusCreated
		commons.DisplayMessage(w, m)
		return
	} else if currentVote.Value != vote.Value {
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		err := updateCommentVotes(vote.CommentID, vote.Value,true)
		if err != nil {
			m.Message = err.Error()
			m.CodState = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "voto actualizado"
		m.CodState = http.StatusOK
		commons.DisplayMessage(w, m)
		return
	}
	m.Message = "Este voto ya esta registrado"
	m.CodState = http.StatusBadRequest
	commons.DisplayMessage(w, m)
}
// actualiza la cantidad de votos en la tabala comentarios
// isUpdate indica si es un voto paraactualizar
func updateCommentVotes(commentID uint, vote bool, isUpdate bool) (err error) {
	comment := models.Comment{}
	db := configuration.GetConnection()
	defer db.Close()
	rows := db.First(&comment, commentID).RowsAffected
	if rows > 0 {
		if vote {
			comment.Votes++
			if isUpdate{
				comment.Votes++
			}
		} else {
			comment.Votes--
			if isUpdate {
				comment.Votes--
			}
		}
		db.Save(&comment)
	} else {
		err = errors.New("No se encontro unregistro de comentario para asignarle el voto")
	}
	return
}
