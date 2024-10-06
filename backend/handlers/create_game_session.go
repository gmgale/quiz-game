package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// PostGames creates a new game session
func PostGames(ctx echo.Context, gameSessions map[string]*models.GameSession) error {
	gameID := uuid.New().String()

	gameSessions[gameID] = &models.GameSession{
		ID:      gameID,
		Status:  "waiting",
		Players: make(map[string]*models.Player),
		Answers: make(map[string][]*models.Answer),
		Questions: []*models.Question{
			{
				ID:        "q1",
				Text:      "What is 2+2?",
				Options:   []string{"2", "3", "4"},
				Answer:    2, // Correct option index
				TimeLimit: 10,
			},
			// Add more questions as needed
		},
		StartTime: time.Time{},
	}

	return ctx.JSON(http.StatusCreated, api.GameSession{
		Id:     ptrString(gameID),
		Status: ptrGameSessionStatus(api.Waiting),
	})
}
