package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// PostGamesGameIdStart starts the game session
func PostGamesGameIdStart(ctx echo.Context, gameId string, gameSessions map[string]*models.GameSession) error {
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

	// TODO: Broadcast the first question to all connected clients (implement WebSocket broadcasting)

	return ctx.JSON(http.StatusOK, api.GameSession{
		Id:     ptrString(gameSession.ID),
		Status: ptrGameSessionStatus(api.Active),
	})
}
