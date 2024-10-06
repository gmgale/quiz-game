package handlers

import (
	"sort"
	"time"

	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
)

func advanceQuestion(gameID string, gameSessions map[string]*models.GameSession) {
	gameSession := gameSessions[gameID]
	question := gameSession.CurrentQuestion

	// Wait for the duration of the question's time limit
	time.Sleep(time.Duration(question.TimeLimit) * time.Second)

	// Move to the next question
	gameSession.CurrentQuestionIndex++

	if gameSession.CurrentQuestionIndex >= len(gameSession.Questions) {
		// Game over
		gameSession.Status = "finished"

		// Prepare and broadcast the leaderboard
		var leaderboard []api.LeaderboardEntry
		for _, player := range gameSession.Players {
			entry := api.LeaderboardEntry{
				PlayerId: ptrString(player.ID),
				Name:     ptrString(player.Name),
				Score:    ptrInt(player.Score),
			}
			leaderboard = append(leaderboard, entry)
		}

		// Sort the leaderboard
		sort.Slice(leaderboard, func(i, j int) bool {
			return *leaderboard[i].Score > *leaderboard[j].Score
		})

		// Broadcast game over message
		msg := models.Message{
			GameID: gameID,
			Type:   "game_over",
			Data:   leaderboard,
		}
		Broadcast <- msg

		return
	}

	// Set the next question
	gameSession.CurrentQuestion = gameSession.Questions[gameSession.CurrentQuestionIndex]
	gameSession.QuestionStartTime = time.Now()

	// Broadcast the next question
	question = gameSession.CurrentQuestion

	msg := models.Message{
		GameID: gameID,
		Type:   "question",
		Data: api.Question{
			Id:        ptrString(question.ID),
			Text:      ptrString(question.Text),
			Options:   &question.Options,
			TimeLimit: ptrInt(question.TimeLimit),
		},
	}

	Broadcast <- msg

	// Start timer for the next question
	go advanceQuestion(gameID, gameSessions)
}
