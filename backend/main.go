package main

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"github.com/gmgale/quiz-game/backend/handlers"
	"github.com/gmgale/quiz-game/backend/models"
	gameServer "github.com/gmgale/quiz-game/backend/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()

	// CORS Middleware configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	server := &gameServer.Server{
		GameSessions: make(map[string]*models.GameSession),
	}

	// Register the handlers
	api.RegisterHandlers(e, server)

	// log any incoming requests
	e.Use(middleware.Logger())

	// Start the message handling goroutine
	go handlers.HandleMessages()

	// Register the WebSocket endpoint
	e.GET("/ws/:gameId", handlers.WebSocketHandler(server.GameSessions))

	// Start the server
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
