package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetGamesGameIdCurrentQuestion retrieves the current question
func (s *Server) GetGamesGameIdCurrentQuestion(ctx echo.Context, gameId string) error {
	gameSession, exists := gameSessions[gameId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	if gameSession.Status != "active" {
		return ctx.JSON(http.StatusBadRequest, "Game is not active")
	}

	question := gameSession.CurrentQuestion

	return ctx.JSON(http.StatusOK, api.Question{
		Id:        &question.ID,
		Text:      &question.Text,
		Options:   &question.Options,
		TimeLimit: &question.TimeLimit,
	})
}
