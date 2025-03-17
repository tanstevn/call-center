package repositories

import (
	"context"

	"call-center/internal/domain/entities"

	"github.com/google/uuid"
)

type CallRepository struct{}

func NewCallRepository() *CallRepository {
	return &CallRepository{}
}

func (callRepo *CallRepository) CreateCall(context context.Context, call *entities.Call) (*entities.Call, error) {
	call.Id = uuid.New()
	return call, nil
}
