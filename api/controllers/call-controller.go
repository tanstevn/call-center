package controllers

import (
	"fmt"
	"net/http"
)

type CallController struct{}

func NewCallController() *CallController {
	fmt.Println("CallController - NewCallController()")
	return &CallController{}
}

func (cc *CallController) CreateCall() http.HandlerFunc {
	fmt.Println("CallController - CreateCall()")

	// call mediator here

	return nil
}

func (cc *CallController) GetCallList() http.HandlerFunc {
	fmt.Println("CallController - GetCallList()")

	// call mediator here

	return nil
}
