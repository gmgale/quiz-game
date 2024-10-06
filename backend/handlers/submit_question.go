package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/labstack/echo/v4"
	"net/http"
)

// PostGamesGameIdAnswers handles answer submission
func (s *Server) PostGamesGameIdAnswers(ctx echo.Context, gameId string) error {
	var answer api.Answer
	if err := ctx.Bind(&answer); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	gameSession, exists := gameSessions[gameId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	player, playerExists := gameSession.Players[*answer.PlayerId]
	if !playerExists {
		return ctx.JSON(http.StatusNotFound, "Player not found")
	}

	// Record the answer
	ans := &Answer{
		PlayerID:       *answer.PlayerId,
		QuestionID:     *answer.QuestionId,
		SelectedOption: int(*answer.SelectedOption),
		ResponseTime:   int64(*answer.ResponseTime),
	}
	gameSession.Answers[player.ID] = append(gameSession.Answers[player.ID], ans)

	// Check if the answer is correct
	currentQuestion := gameSession.CurrentQuestion
	correct := currentQuestion.Answer == ans.SelectedOption

	// Calculate score (simple example: 10 points for correct answer)
	scoreAwarded := 0
	if correct {
		scoreAwarded = 10
		player.Score += scoreAwarded
	}

	return ctx.JSON(http.StatusOK, api.AnswerResponse{
		Correct:       &correct,
		CorrectOption: &currentQuestion.Answer,
		ScoreAwarded:  &scoreAwarded,
	})
}
