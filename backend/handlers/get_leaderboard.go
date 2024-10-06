package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sort"
)

// GetGamesGameIdLeaderboard retrieves the leaderboard
func GetGamesGameIdLeaderboard(ctx echo.Context, gameId string, gameSessions map[string]*models.GameSession) error {
	gameSession, exists := gameSessions[gameId]
	if !exists {
		return ctx.JSON(http.StatusNotFound, "Game session not found")
	}

	var leaderboard []api.LeaderboardEntry
	for _, player := range gameSession.Players {
		entry := api.LeaderboardEntry{
			PlayerId: ptrString(player.ID),
			Name:     ptrString(player.Name),
			Score:    ptrInt(player.Score),
		}
		leaderboard = append(leaderboard, entry)
	}

	// Sort the leaderboard by score
	sort.Slice(leaderboard, func(i, j int) bool {
		return *leaderboard[i].Score > *leaderboard[j].Score
	})

	return ctx.JSON(http.StatusOK, leaderboard)
}
