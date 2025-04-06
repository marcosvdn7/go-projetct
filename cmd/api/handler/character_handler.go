package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcosvdn7/go-projetct/cmd/api/request"
	"net/http"
)

func CreateCharacterHandler(ctx *gin.Context) {
	var character *request.CharacterRequest

	logger.Infof("Create character %v", character)
	if err := ctx.BindJSON(&character); err != nil {
		logger.Errorf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading request body: " + err.Error(),
		})
		return
	}

	response, err := request.CreateCharacter(character)
	if err != nil {
		logger.Errorf("Error creating character: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error creating character: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func GetCharacterHandler(ctx *gin.Context) {
	id := parseUUIDFromQueryParam(ctx)

	character, err := request.GetCharacter(id)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": character,
	})
}

func UpdateCharacterHandler(ctx *gin.Context) {
	id := parseUUIDFromQueryParam(ctx)

	var character *request.CharacterRequest
	if err := ctx.BindJSON(&character); err != nil {
		logger.Errorf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error reading request body: " + err.Error(),
		})
		return
	}

	response, err := request.UpdateCharacter(id, character)
	if err != nil {
		logger.Errorf("Error updating character: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error updating character: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func DeleteCharacterHandler(ctx *gin.Context) {
	id := parseUUIDFromQueryParam(ctx)
	rowsAffected, err := request.DeleteCharacter(id)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error deleting character: %v", err),
		})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Character with id %v not found to delete", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Character with id %v succesfully deleted", id),
	})
}

func ListCharactersHandler(ctx *gin.Context) {
	result, err := request.ListCharacters()
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func parseUUIDFromQueryParam(ctx *gin.Context) uuid.UUID {
	stringId := ctx.Query("id")
	if stringId == "" {
		logger.Errorf("Query param id is required")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Query param id is required",
		})
		return uuid.Nil
	}
	id, err := uuid.Parse(stringId)
	if err != nil {
		logger.Errorf("Error parsing uuid: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Invalid uuid: %s", stringId),
		})
		return uuid.Nil
	}

	return id
}
