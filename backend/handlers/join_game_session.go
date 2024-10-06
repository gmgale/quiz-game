package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// PostGamesGameIdPlayers handles a player joining a game session
func (s *Server) PostGamesGameIdPlayers(ctx echo.Context, gameId string) error {
	var req struct {
		Name string `json:"name"`
	}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	gameSession, exists := gameSessions[gameId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	playerID := uuid.New().String()
	player := &Player{
		ID:       playerID,
		Name:     req.Name,
		Score:    0,
		JoinedAt: time.Now(),
	}
	gameSession.Players[playerID] = player

	return ctx.JSON(http.StatusOK, api.Player{
		Id:       &player.ID,
		Name:     &player.Name,
		Score:    &player.Score,
		JoinedAt: ptrTime(player.JoinedAt),
	})
}
