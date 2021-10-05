package models

type Comment struct {
	Message string `form:"message" json:"message"`
}
