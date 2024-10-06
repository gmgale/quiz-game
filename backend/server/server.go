package server

import (
	"github.com/gmgale/quiz-game/backend/handlers"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/labstack/echo/v4"
	"sync"
)

type Server struct {
	GameSessions map[string]*models.GameSession
	Mutex        sync.Mutex
}

// PostGames handles the creation of a new game session
func (s *Server) PostGames(ctx echo.Context) error {
	return handlers.PostGames(ctx, s.GameSessions)
}

// PostGamesGameIdPlayers handles a player joining a game session
func (s *Server) PostGamesGameIdPlayers(ctx echo.Context, gameId string) error {
	return handlers.PostGamesGameIdPlayers(ctx, gameId, s)
}

// PostGamesGameIdStart starts the game session
func (s *Server) PostGamesGameIdStart(ctx echo.Context, gameId string) error {
	return handlers.PostGamesGameIdStart(ctx, gameId, s.GameSessions)
}

// GetGamesGameIdLeaderboard retrieves the leaderboard
func (s *Server) GetGamesGameIdLeaderboard(ctx echo.Context, gameId string) error {
	return handlers.GetGamesGameIdLeaderboard(ctx, gameId, s.GameSessions)
}

// PostGamesGameIdAnswers handles answer submission
func (s *Server) PostGamesGameIdAnswers(ctx echo.Context, gameId string) error {
	return handlers.PostGamesGameIdAnswers(ctx, gameId, s.GameSessions)
}

// GetGamesGameIdCurrentQuestion retrieves the current question
func (s *Server) GetGamesGameIdCurrentQuestion(ctx echo.Context, gameId string) error {
	return handlers.GetGamesGameIdCurrentQuestion(ctx, gameId, s.GameSessions)
}
