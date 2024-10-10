package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/gmgale/quiz-game/backend/questions"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

// PostGames creates a new game session
func PostGames(ctx echo.Context, gameSessions map[string]*models.GameSession, loadedQuestions []questions.Question) error {
	gameID := uuid.New().String()
	// Generate a unique code for the game session
	gameCode := uuid.New().String()[:6] // short code example, adjust as needed
	log.Print("Generated game code: ", gameCode)

	// Convert loaded questions to models.Question type
	var gameQuestions []*models.Question
	for _, q := range loadedQuestions {
		gameQuestions = append(gameQuestions, &models.Question{
			ID:        q.ID,
			Text:      q.Text,
			Options:   q.Options,
			TimeLimit: q.TimeLimit,
			Answer:    -1, // Placeholder, correct answer not exposed here
		})
	}

	gameSessions[gameID] = &models.GameSession{
		ID:        gameID,
		Code:      gameCode,
		Status:    "waiting",
		Players:   make(map[string]*models.Player),
		Answers:   make(map[string][]*models.Answer),
		Questions: gameQuestions,
		StartTime: time.Time{},
	}

	return ctx.JSON(http.StatusCreated, api.GameSession{
		Id:     ptrString(gameID),
		Status: ptrGameSessionStatus(api.Waiting),
		Code:   ptrString(gameCode),
	})
}
