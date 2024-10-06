package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

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
	gameSession.QuestionStartTime = time.Now()

	// Broadcast the first question
	question := gameSession.CurrentQuestion

	msg := models.Message{
		GameID: gameId,
		Type:   "question",
		Data: api.Question{
			Id:        ptrString(question.ID),
			Text:      ptrString(question.Text),
			Options:   &question.Options,
			TimeLimit: ptrInt(question.TimeLimit),
		},
	}

	Broadcast <- msg

	// Start the timer for the question
	go advanceQuestion(gameId, gameSessions)

	return ctx.JSON(http.StatusOK, api.GameSession{
		Id:     ptrString(gameSession.ID),
		Status: ptrGameSessionStatus(api.Active),
	})
}
