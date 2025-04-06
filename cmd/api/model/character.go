package model

import "github.com/google/uuid"

type Character struct {
	Id         uuid.UUID
	Name       string
	Level      int
	Class      string
	Specie     string
	Initiative int
	Speed      int
	HP         int
}
