package main

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/handlers"
	"github.com/gmgale/quiz-game/backend/models"
	gameServer "github.com/gmgale/quiz-game/backend/server"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()

	server := &gameServer.Server{
		GameSessions: make(map[string]*models.GameSession),
	}

	// Register the handlers
	api.RegisterHandlers(e, server)

	// Start the message handling goroutine
	go handlers.HandleMessages()

	// Register the WebSocket endpoint
	e.GET("/ws/:gameId", handlers.WebSocketHandler(server.GameSessions))

	// Start the server
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
