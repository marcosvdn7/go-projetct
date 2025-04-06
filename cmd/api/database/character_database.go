package database

import (
	"errors"
	"github.com/google/uuid"
	"github.com/marcosvdn7/go-projetct/cmd/api/model"
	"strings"
)

func SaveCharacter(c *model.Character) error {
	return db.QueryRow("INSERT INTO \"go-project\".character(name, class, specie, initiative, speed, hp, level) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, class, specie, initiative, speed, hp, level",
		c.Name, c.Class, c.Specie, c.Initiative, c.Speed, c.HP, c.Level).Scan(
		&c.Id,
		&c.Name,
		&c.Class,
		&c.Specie,
		&c.Initiative,
		&c.Speed,
		&c.HP,
		&c.Level,
	)
}

func FindCharacterById(id uuid.UUID) (c *model.Character, err error) {
	c = &model.Character{}
	err = db.QueryRow("SELECT id, name, class, specie, initiative, speed, hp, level "+
		"FROM \"go-project\".character WHERE id = $1", id).
		Scan(
			&c.Id,
			&c.Name,
			&c.Class,
			&c.Specie,
			&c.Initiative,
			&c.Speed,
			&c.HP,
			&c.Level,
		)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			logger.Errorf("id not found")
			return nil, errors.New("character not found")
		} else {
			logger.Errorf("Error while fetching character: %v", err)
			return nil, err
		}
	}

	return c, nil
}

func UpdateCharacter(id uuid.UUID, c *model.Character) (err error) {
	return db.QueryRow("UPDATE \"go-project\".character "+
		"SET name = $1, class = $2, specie = $3, initiative = $4, speed = $5, hp = $6, level = $7 "+
		"WHERE id = $8 "+
		"RETURNING id, name, class, specie, initiative, speed, hp, level",
		c.Name, c.Class, c.Specie, c.Initiative, c.Speed, c.HP, c.Level, id).
		Scan(
			&c.Id,
			&c.Name,
			&c.Class,
			&c.Specie,
			&c.Initiative,
			&c.Speed,
			&c.HP,
			&c.Level,
		)
}

func DeleteCharacter(id uuid.UUID) (rowsAffected int64, err error) {
	result, err := db.Exec("DELETE FROM \"go-project\".character WHERE id = $1", id)
	if err != nil {
		logger.Errorf("Error while deleting character: %v", err)
		return 0, err
	}
	return result.RowsAffected()
}

func ListCharacters() (characters []*model.Character, err error) {
	result, err := db.Query("SELECT id, name, class, specie, initiative, speed, hp, level " +
		"FROM \"go-project\".character")
	if err != nil {
		logger.Errorf("Error while fetching characters: %v", err)
		return nil, err
	}

	defer result.Close()

	character := &model.Character{}
	for result.Next() {
		err = result.Scan(
			&character.Id,
			&character.Name,
			&character.Class,
			&character.Specie,
			&character.Initiative,
			&character.Speed,
			&character.HP,
			&character.Level,
		)
		if err != nil {
			logger.Errorf("Error while fetching character: %v", err)
			return characters, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}
