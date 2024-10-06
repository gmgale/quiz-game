package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/gmgale/quiz-game/backend/server"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// PostGamesGameIdPlayers handles a player joining a game session
func PostGamesGameIdPlayers(ctx echo.Context, gameId string, server *server.Server) error {
	server.Mutex.Lock()
	gameSession, exists := server.GameSessions[gameId]
	server.Mutex.Unlock()
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	server.Mutex.Lock()
	defer server.Mutex.Unlock()
	var req struct {
		Name string `json:"name"`
	}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	gameSession, exists = server.GameSessions[gameId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	playerID := uuid.New().String()
	player := &models.Player{
		ID:       playerID,
		Name:     req.Name,
		Score:    0,
		JoinedAt: time.Now(),
	}
	gameSession.Players[playerID] = player

	return ctx.JSON(http.StatusOK, api.Player{
		Id:       ptrString(player.ID),
		Name:     ptrString(player.Name),
		Score:    ptrInt(player.Score),
		JoinedAt: ptrTime(player.JoinedAt),
	})
}
