package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"time"
)

func ptrString(s string) *string                                          { return &s }
func ptrGameSessionStatus(s api.GameSessionStatus) *api.GameSessionStatus { return &s }
func ptrTime(t time.Time) *time.Time                                      { return &t }
