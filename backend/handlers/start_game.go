package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// PostGamesGameIdStart starts the game session
func (s *Server) PostGamesGameIdStart(ctx echo.Context, gameId string) error {
	gameSession, exists := gameSessions[gameId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	if gameSession.Status != "waiting" {
		return ctx.JSON(http.StatusBadRequest, "Game has already started or finished")
	}

	gameSession.Status = "active"
	gameSession.StartTime = time.Now()
	gameSession.CurrentQuestionIndex = 0
	gameSession.CurrentQuestion = gameSession.Questions[0]

	// Broadcast the first question to all connected clients (implement WebSocket broadcasting)

	return ctx.JSON(http.StatusOK, api.GameSession{
		Id:     &gameSession.ID,
		Status: ptrGameSessionStatus(api.GameSessionStatus(gameSession.Status)),
	})
}
