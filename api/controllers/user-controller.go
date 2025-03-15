package controllers

import (
	"fmt"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUserList() http.HandlerFunc {
	fmt.Println("UserController - GetUserList()")

	// call mediator here

	return nil
}
