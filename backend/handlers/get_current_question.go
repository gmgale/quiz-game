package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetGamesGameIdCurrentQuestion retrieves the current question
func GetGamesGameIdCurrentQuestion(ctx echo.Context, gameId string, gameSessions map[string]*models.GameSession) error {
	gameSession, exists := gameSessions[gameId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	if gameSession.Status != "active" {
		return ctx.JSON(http.StatusBadRequest, "Game is not active")
	}

	question := gameSession.CurrentQuestion

	return ctx.JSON(http.StatusOK, api.Question{
		Id:        ptrString(question.ID),
		Text:      ptrString(question.Text),
		Options:   &question.Options,
		TimeLimit: ptrInt(question.TimeLimit),
	})
}
