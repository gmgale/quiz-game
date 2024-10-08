package server

import (
	"github.com/gmgale/quiz-game/backend/handlers"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/gmgale/quiz-game/backend/questions"
	"github.com/labstack/echo/v4"
	"log"
	"sync"
)

type Server struct {
	GameSessions map[string]*models.GameSession
	Mutex        sync.Mutex
	Questions    []questions.Question
}

func NewServer(questionsFile string) *Server {
	loadedQuestions, err := questions.LoadQuestions(questionsFile)
	if err != nil {
		log.Fatalf("Failed to load questions: %v", err)
	}
	return &Server{
		GameSessions: make(map[string]*models.GameSession),
		Questions:    loadedQuestions,
	}
}

// PostGames handles the creation of a new game session
func (s *Server) PostGames(ctx echo.Context) error {
	return handlers.PostGames(ctx, s.GameSessions, s.Questions)
}

func (s *Server) PostGamesGameIdPlayers(ctx echo.Context, gameId string) error {
	return handlers.PostGamesGameIdPlayers(ctx, s.GameSessions, &s.Mutex)
}

func (s *Server) PostGamesGameIdPlayersByID(ctx echo.Context, gameId string) error {
	return s.PostGamesGameIdPlayers(ctx, gameId)
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
