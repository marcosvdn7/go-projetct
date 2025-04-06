package request

import (
	"errors"
	"github.com/google/uuid"
	"github.com/marcosvdn7/go-projetct/cmd/api/database"
	"github.com/marcosvdn7/go-projetct/cmd/api/model"
)

type CharacterRequest struct {
	Name       string `json:"name"`
	Class      string `json:"class"`
	Level      *int   `json:"level"`
	Specie     string `json:"specie"`
	Initiative *int   `json:"initiative"`
	Speed      *int   `json:"speed"`
	HP         *int   `json:"hp"`
}

type CharacterResponse struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Class      string    `json:"class"`
	Level      int       `json:"level"`
	Specie     string    `json:"specie"`
	Initiative int       `json:"initiative"`
	Speed      int       `json:"speed"`
	HP         int       `json:"hp"`
}

func CreateCharacter(character *CharacterRequest) (c *CharacterResponse, err error) {
	if err = character.validateCharacterCreate(); err != nil {
		return nil, err
	}

	characterToPersist := mapRequestToModel(character)

	if err = database.SaveCharacter(characterToPersist); err != nil {
		return nil, err
	}

	response := mapModelToResponse(characterToPersist)

	return response, nil
}

func GetCharacter(id uuid.UUID) (c *CharacterResponse, err error) {
	character, err := database.FindCharacterById(id)
	if err != nil {
		return nil, err
	}
	response := mapModelToResponse(character)

	return response, err
}

func UpdateCharacter(id uuid.UUID, character *CharacterRequest) (c *CharacterResponse, err error) {
	persistedCharacter, err := database.FindCharacterById(id)
	if err != nil {
		return nil, err
	}

	mapFieldsToUpdate(persistedCharacter, character)

	err = database.UpdateCharacter(id, persistedCharacter)
	if err != nil {
		return nil, err
	}

	response := mapModelToResponse(persistedCharacter)

	return response, nil
}

func DeleteCharacter(id uuid.UUID) (rowsAffected int64, err error) {
	return database.DeleteCharacter(id)
}

func ListCharacters() (characters []*CharacterResponse, err error) {
	result, err := database.ListCharacters()
	if err != nil {
		return nil, err
	}

	for _, res := range result {
		character := mapModelToResponse(res)
		characters = append(characters, character)
	}

	return characters, nil
}

func (c *CharacterRequest) validateCharacterCreate() error {
	err := validateRequiredFields(c)

	if err != nil {
		return err
	}

	return nil
}

func validateRequiredFields(c *CharacterRequest) error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Class == "" {
		return errors.New("class is required")
	}
	if c.Specie == "" {
		return errors.New("specie is required")
	}
	if c.Initiative == nil {
		return errors.New("initiative is required")
	}
	if c.Speed == nil {
		return errors.New("speed is required")
	}
	if c.HP == nil {
		return errors.New("hp is required")
	}

	if c.Level == nil {
		return errors.New("level is required")
	}

	return nil
}

func mapRequestToModel(c *CharacterRequest) *model.Character {
	return &model.Character{
		Name:       c.Name,
		Class:      c.Class,
		Specie:     c.Specie,
		Level:      *c.Level,
		Initiative: *c.Initiative,
		Speed:      *c.Speed,
		HP:         *c.HP,
	}
}

func mapModelToResponse(character *model.Character) *CharacterResponse {
	return &CharacterResponse{
		Id:         character.Id,
		Name:       character.Name,
		Class:      character.Class,
		Level:      character.Level,
		Specie:     character.Specie,
		Initiative: character.Initiative,
		Speed:      character.Speed,
		HP:         character.HP,
	}
}

func mapFieldsToUpdate(persisted *model.Character, updated *CharacterRequest) {
	if updated.Name != "" {
		persisted.Name = updated.Name
	}
	if updated.Class != "" {
		persisted.Class = updated.Class
	}
	if updated.Specie != "" {
		persisted.Specie = updated.Specie
	}
	if updated.Level != nil {
		persisted.Level = *updated.Level
	}
	if updated.Initiative != nil {
		persisted.Initiative = *updated.Initiative
	}
	if updated.Speed != nil {
		persisted.Speed = *updated.Speed
	}
	if updated.HP != nil {
		persisted.HP = *updated.HP
	}
}
