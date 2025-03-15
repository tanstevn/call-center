package repositories

import (
	"context"

	"call-center/internal/domain/entities"
)

type CallRepository struct{}

func NewCallRepository() *CallRepository {
	return &CallRepository{}
}

func (callRepo *CallRepository) CreateCall(context context.Context, call *entities.Call) (*entities.Call, error) {
	return &entities.Call{}, nil
}
