package model

import "github.com/google/uuid"

type Character struct {
	Id         uuid.UUID
	Name       string
	Level      int
	Class      string
	Specie     Specie
	Initiative int
	Speed      int
	HP         int
	Attributes Attribute
}
