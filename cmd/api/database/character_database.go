package database

import (
	"errors"
	"github.com/google/uuid"
	"github.com/marcosvdn7/go-projetct/cmd/api/model"
	"strings"
)

const (
	schemaName = "\"go-project\""
)

func SaveCharacter(c *model.Character) error {
	return db.QueryRow("INSERT INTO \"go-project\".character(name, class, specie, initiative, speed, hp, level) "+
		"VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, class, initiative, speed, hp, level, "+
		"(SELECT s.name FROM \"go-project\".specie as s WHERE s.id = $8) as specie_name",
		c.Name, c.Class, c.Specie.Id, c.Initiative, c.Speed, c.HP, c.Level, c.Specie.Id).Scan(
		&c.Id,
		&c.Name,
		&c.Class,
		&c.Initiative,
		&c.Speed,
		&c.HP,
		&c.Level,
		&c.Specie.Name,
	)
}

func FindCharacterById(id uuid.UUID) (c *model.Character, err error) {
	c = &model.Character{}
	err = db.QueryRow("SELECT c.id, c.name, c.class, s.name, c.initiative, c.speed, c.hp, c.level "+
		"FROM \"go-project\".character as c "+
		"LEFT JOIN \"go-project\".specie as s ON s.id = c.specie "+
		"WHERE c.id = $1", id).
		Scan(
			&c.Id,
			&c.Name,
			&c.Class,
			&c.Specie.Name,
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
		"RETURNING id, name, class, initiative, speed, hp, level, "+
		"(SELECT s.name FROM \"go-project\".specie as s WHERE s.id = $9) as specie_name",
		c.Name, c.Class, c.Specie.Id, c.Initiative, c.Speed, c.HP, c.Level, id, c.Specie.Id).
		Scan(
			&c.Id,
			&c.Name,
			&c.Class,
			&c.Initiative,
			&c.Speed,
			&c.HP,
			&c.Level,
			&c.Specie.Name,
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
	result, err := db.Query("SELECT c.id, c.name, c.class, s.name, c.initiative, c.speed, c.hp, c.level " +
		"FROM \"go-project\".character AS c " +
		"LEFT JOIN \"go-project\".specie AS s on s.id = c.specie")
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
			&character.Specie.Name,
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
