package handlers

import (
	"github.com/gmgale/quiz-game/backend/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sync"
)

// Upgrader config
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust this in production for security
	},
}

// Client represents a WebSocket client
type Client struct {
	Conn     *websocket.Conn
	GameCode string
	Mutex    sync.Mutex
}

// Map of clients per game session
var clients = make(map[string]map[*Client]bool)
var clientsMutex sync.RWMutex

// Broadcast channel
var Broadcast = make(chan models.Message)

// WebSocketHandler handles WebSocket requests from clients
func WebSocketHandler(gameSessions map[string]*models.GameSession) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("WebSocket connection from %s", c.RealIP())
		gameCode := c.Param("gameCode")
		_, exists := gameSessions[gameCode]
		if !exists {
			return c.JSON(http.StatusNotFound, "Game session not found")
		}

		// Upgrade the connection to WebSocket
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Printf("WebSocket upgrade error: %v", err)
			return err
		}

		client := &Client{Conn: ws, GameCode: gameCode}

		// Register the client
		clientsMutex.Lock()
		if clients[gameCode] == nil {
			clients[gameCode] = make(map[*Client]bool)
		}
		clients[gameCode][client] = true
		clientsMutex.Unlock()

		// Start listening for messages from the client
		go client.readMessages(gameSessions)

		return nil
	}
}

// readMessages listens for incoming messages from the client
func (client *Client) readMessages(gameSessions map[string]*models.GameSession) {
	defer func() {
		client.Conn.Close()
		clientsMutex.Lock()
		delete(clients[client.GameCode], client)
		clientsMutex.Unlock()
	}()

	for {
		var msg models.Message
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}
		// Handle received message if needed
	}
}

// HandleMessages broadcasts messages to all clients in a game session
func HandleMessages() {
	for {
		msg := <-Broadcast
		clientsMutex.RLock()
		for client := range clients[msg.GameCode] {
			client.Mutex.Lock()
			err := client.Conn.WriteJSON(msg)
			client.Mutex.Unlock()
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				client.Conn.Close()
				clientsMutex.Lock()
				delete(clients[msg.GameCode], client)
				clientsMutex.Unlock()
			}
		}
		clientsMutex.RUnlock()
	}
}
