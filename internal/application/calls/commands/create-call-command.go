package commands

import (
	"call-center/internal/domain/entities"
	"call-center/internal/infrastructure/data/repositories"
	"context"
	"fmt"
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

func NewCreateCallCommand(callingUserId uuid.UUID, dateCallStarted time.Time) *CreateCallCommand {
	return &CreateCallCommand{
		CallingUserId:   callingUserId,
		DateCallStarted: dateCallStarted,
	}
}

func NewCreateCallCommandHandler(repository *repositories.CallRepository) *CreateCallCommandHandler {
	return &CreateCallCommandHandler{repository: repository}
}

func (handler *CreateCallCommandHandler) Handle(ctx context.Context, request *CreateCallCommand) (*CreateCallResult, error) {
	isLoggerPipelineEnabled := ctx.Value("logger_pipeline").(bool)

	if isLoggerPipelineEnabled {
		fmt.Printf("[%T]: logging pipeline is enabled", handler)
	}

	call := entities.Call{}
	response, err := handler.repository.CreateCall(ctx, &call)

	if err != nil {
		return nil, err
	}

	result := CreateCallResult{
		Id: response.Id,
	}

	return &result, nil
}
