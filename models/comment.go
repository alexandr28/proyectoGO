package models

import "github.com/jinzhu/gorm"

// Comment del sistema
type Comment struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	ParentID uint      `json:"parent_id"`
	Votes    int32     `json:"votes"`
	Content  string    `json:"content"`
	HasVote  int8      `json:"hasvote" gorm:"-"`
	User     []User    `json:"user,omitempty"`
	Children []Comment `json:"children,omitempty"`
}
