package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sync"
	"time"
)

// PostGamesGameIdPlayers handles a player joining a game session by code
func PostGamesGameIdPlayers(ctx echo.Context, gameSessions map[string]*models.GameSession, mutex *sync.Mutex) error {
	var req struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}

	// Bind the request body to the struct
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Find the game session by code (not by ID)
	log.Print("Joining game session with code: ", req)
	var gameSession *models.GameSession
	mutex.Lock()
	for _, session := range gameSessions {
		if session.Code == req.Code {
			gameSession = session
			break
		}
	}
	mutex.Unlock()

	if gameSession == nil {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	// Create a new player and add them to the session
	playerID := uuid.New().String()
	player := &models.Player{
		ID:       playerID,
		Name:     req.Name,
		Score:    0,
		JoinedAt: time.Now(),
	}
	mutex.Lock()
	gameSession.Players[playerID] = player
	mutex.Unlock()

	return ctx.JSON(http.StatusOK, api.Player{
		Id:       ptrString(player.ID),
		Name:     ptrString(player.Name),
		Score:    ptrInt(player.Score),
		JoinedAt: ptrTime(player.JoinedAt),
	})
}
