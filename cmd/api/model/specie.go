package model

import "github.com/google/uuid"

type Specie struct {
	Id                   uuid.UUID
	Name                 string
	AbilityScoreIncrease string
	Size                 string
	Speed                int
}
