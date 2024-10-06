package main

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// Server implements the ServerInterface from api.gen.go
type Server struct{}

// PostGames handles the creation of a new game session
func (s *Server) PostGames(ctx echo.Context) error {
	// Implement your logic here
	return ctx.JSON(http.StatusCreated, api.GameSession{
		Id:     ptrString("game-session-id"),
		Status: ptrGameSessionStatus(api.Waiting),
	})
}

// PostGamesGameIdPlayers handles a player joining a game session
func (s *Server) PostGamesGameIdPlayers(ctx echo.Context, gameId string) error {
	// Implement your logic here
	return ctx.JSON(http.StatusOK, api.Player{
		Id:   ptrString("player-id"),
		Name: ptrString("Player Name"),
	})
}

func main() {
	e := echo.New()

	// Create an instance of your server
	server := &Server{}

	// Register the handlers from the generated code
	api.RegisterHandlers(e, server)

	// Start the server
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Helper functions for pointer types
func ptrString(s string) *string { return &s }
func ptrGameSessionStatus(s api.GameSessionStatus) *api.GameSessionStatus {
	return &s
}
