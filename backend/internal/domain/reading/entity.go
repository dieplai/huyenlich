package reading

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ServiceType string

const (
	ServiceTarot ServiceType = "TAROT"
)

type Reading struct {
	ID          uuid.UUID
	UserID      *uuid.UUID
	ServiceType ServiceType
	InputParams json.RawMessage
	FreeResult  json.RawMessage
	PaidResult  *json.RawMessage
	AIContent   *string
	IsUnlocked  bool
	CreatedAt   time.Time
}

func New(serviceType ServiceType, inputParams json.RawMessage, freeResult json.RawMessage) *Reading {
	return &Reading{
		ID:          uuid.New(),
		ServiceType: serviceType,
		InputParams: inputParams,
		FreeResult:  freeResult,
		IsUnlocked:  false,
		CreatedAt:   time.Now(),
	}
}

func (r *Reading) Unlock(paidResult json.RawMessage, aiContent string) {
	r.PaidResult = &paidResult
	r.AIContent = &aiContent
	r.IsUnlocked = true
}
