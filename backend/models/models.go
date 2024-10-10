package models

import (
	"time"
)

// GameSession represents a game session
type GameSession struct {
	ID                   string
	Code                 string
	Status               string
	CurrentQuestion      *Question
	StartTime            time.Time
	Players              map[string]*Player
	Questions            []*Question
	Answers              map[string][]*Answer
	CurrentQuestionIndex int
	QuestionStartTime    time.Time
}

// Player represents a player in the game
type Player struct {
	ID       string
	Name     string
	Score    int
	JoinedAt time.Time
}

// Question represents a quiz question
type Question struct {
	ID        string
	Text      string
	Options   []string
	TimeLimit int
	Answer    int // Correct option index
}

// Answer represents a player's answer to a question
type Answer struct {
	PlayerID       string
	QuestionID     string
	SelectedOption int
	ResponseTime   int64 // in milliseconds
}

// LeaderboardEntry represents an entry in the leaderboard
type LeaderboardEntry struct {
	PlayerID string
	Name     string
	Score    int
}

type Message struct {
	GameID   string      `json:"gameId"`
	GameCode string      `json:"gameCode"`
	Type     string      `json:"type"`
	Data     interface{} `json:"data"`
}
