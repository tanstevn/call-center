package commands

import (
	"call-center/internal/infrastructure/data/repositories"
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateCallCommand struct {
	CallingUserId   uuid.UUID `validate:"required"`
	DateCallStarted time.Time `validate:"required"`
}

type CreateCallResult struct {
	Id uuid.UUID `json:"id"`
}

type CreateCallCommandHandler struct {
	repository *repositories.CallRepository
}

func NewCreateCallCommandHandler() *CreateCallCommandHandler {
	return &CreateCallCommandHandler{}
}

func (handler *CreateCallCommandHandler) Handle(context context.Context, request *CreateCallCommand) (*CreateCallResult, error) {
	return &CreateCallResult{}, nil
}
