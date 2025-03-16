package v1

import (
	"call-center/internal/application/calls/commands"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	mediator "github.com/mehdihadeli/go-mediatr"
)

type CallController struct{}

func NewCallController() *CallController {
	fmt.Println("CallController - NewCallController()")
	return &CallController{}
}

func (cc *CallController) CreateCall() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("CallController - CreateCall()")

		command := commands.NewCreateCallCommand(uuid.New(), time.Now().UTC())
		result, err := mediator.Send[*commands.CreateCallCommand, commands.CreateCallResult](request.Context(), command)

		if err != nil {

		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)

		json.NewEncoder(writer).Encode(result)
	}
}

func (cc *CallController) GetCallList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("CallController - GetCallList()")

		// call mediator here
	}
}
