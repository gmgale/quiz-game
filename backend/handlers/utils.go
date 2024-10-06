package handlers

import (
	api "github.com/gmgale/quiz-game/backend/gen"
	"time"
)

func ptrString(s string) *string     { return &s }
func ptrInt(i int) *int              { return &i }
func ptrBool(b bool) *bool           { return &b }
func ptrTime(t time.Time) *time.Time { return &t }
func ptrGameSessionStatus(s api.GameSessionStatus) *api.GameSessionStatus {
	return &s
}
