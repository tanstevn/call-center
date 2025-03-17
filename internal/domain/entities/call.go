package entities

import (
	"time"

	"github.com/google/uuid"
)

type Call struct {
	Id              uuid.UUID
	CallingUserId   uuid.UUID
	DateCallStarted time.Time
}
