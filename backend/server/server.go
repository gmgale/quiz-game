package server

import (
	"github.com/gmgale/quiz-game/backend/handlers"
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/gmgale/quiz-game/backend/questions"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"sync"
	"time"
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

	// Create a default game session
	defaultGameID := uuid.New().String()
	defaultGameCode := "123456" // Use a fixed code for testing

	var gameQuestions []*models.Question
	for _, q := range loadedQuestions {
		gameQuestions = append(gameQuestions, &models.Question{
			ID:        q.ID,
			Text:      q.Text,
			Options:   q.Options,
			TimeLimit: q.TimeLimit,
			Answer:    -1, // Placeholder
		})
	}

	defaultGameSession := &models.GameSession{
		ID:        defaultGameID,
		Code:      defaultGameCode,
		Status:    "waiting",
		Players:   make(map[string]*models.Player),
		Answers:   make(map[string][]*models.Answer),
		Questions: gameQuestions,
		StartTime: time.Time{},
	}

	gameSessions := make(map[string]*models.GameSession)
	gameSessions[defaultGameID] = defaultGameSession

	return &Server{
		GameSessions: gameSessions,
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
