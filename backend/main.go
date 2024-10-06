package main

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/handlers"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/labstack/echo/v4"
	"log"
)

type Server struct {
	GameSessions map[string]*models.GameSession
}

// PostGames handles the creation of a new game session
func (s *Server) PostGames(ctx echo.Context) error {
	return handlers.PostGames(ctx, s.GameSessions)
}

// PostGamesGameIdPlayers handles a player joining a game session
func (s *Server) PostGamesGameIdPlayers(ctx echo.Context, gameId string) error {
	return handlers.PostGamesGameIdPlayers(ctx, gameId, s.GameSessions)
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

func main() {
	e := echo.New()

	server := &Server{
		GameSessions: make(map[string]*models.GameSession),
	}

	// Register the handlers
	api.RegisterHandlers(e, server)

	// Start the server
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
