package request

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/marcosvdn7/go-projetct/cmd/api/database"
	"github.com/marcosvdn7/go-projetct/cmd/api/model"
)

type CharacterRequest struct {
	Name       string         `json:"name"`
	Class      string         `json:"class"`
	Level      *int           `json:"level"`
	Specie     *SpecieRequest `json:"specie"`
	Initiative *int           `json:"initiative"`
	Speed      *int           `json:"speed"`
	HP         *int           `json:"hp"`
}

type CharacterResponse struct {
	Id         uuid.UUID       `json:"id"`
	Name       string          `json:"name"`
	Class      string          `json:"class"`
	Level      int             `json:"level"`
	Specie     *SpecieResponse `json:"specie,omitempty"`
	Initiative int             `json:"initiative"`
	Speed      int             `json:"speed"`
	HP         int             `json:"hp"`
}

type SpecieRequest struct {
	Id uuid.UUID `json:"id"`
}

type SpecieResponse struct {
	Name string `json:"name,omitempty"`
}

func CreateCharacter(character *CharacterRequest) (c *CharacterResponse, err error) {
	if err = character.validateCharacterCreate(); err != nil {
		return nil, err
	}

	characterToPersist, err := mapRequestToModel(character)
	if err != nil {
		return nil, err
	}

	if err = database.SaveCharacter(characterToPersist); err != nil {
		return nil, err
	}

	response, err := mapModelToResponse(characterToPersist)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetCharacter(id uuid.UUID) (c *CharacterResponse, err error) {
	character, err := database.FindCharacterById(id)
	if err != nil {
		return nil, err
	}
	response, err := mapModelToResponse(character)
	if err != nil {
		return nil, err
	}

	return response, nil
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

	response, err := mapModelToResponse(persistedCharacter)
	if err != nil {
		return nil, err
	}

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
		character, err := mapModelToResponse(res)
		if err != nil {
			return nil, err
		}
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
	if c.Specie == nil {
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

func mapRequestToModel(c *CharacterRequest) (modelCharacter *model.Character, err error) {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &modelCharacter)
	if err != nil {
		return nil, err
	}

	return modelCharacter, nil
}

func mapModelToResponse(character *model.Character) (characterResponse *CharacterResponse, err error) {
	jsonData, err := json.Marshal(character)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &characterResponse)
	if err != nil {
		return nil, err
	}

	return characterResponse, nil
}

func mapFieldsToUpdate(persisted *model.Character, updated *CharacterRequest) {
	if updated.Name != "" && updated.Name != persisted.Name {
		persisted.Name = updated.Name
	}
	if updated.Class != "" && updated.Class != persisted.Class {
		persisted.Class = updated.Class
	}
	if updated.Specie != nil {
		persisted.Specie.Id = updated.Specie.Id
	}
	if updated.Level != nil && *updated.Level != persisted.Level {
		persisted.Level = *updated.Level
	}
	if updated.Initiative != nil && *updated.Initiative != persisted.Initiative {
		persisted.Initiative = *updated.Initiative
	}
	if updated.Speed != nil && *updated.Speed != persisted.Speed {
		persisted.Speed = *updated.Speed
	}
	if updated.HP != nil && *updated.HP != persisted.HP {
		persisted.HP = *updated.HP
	}
}
