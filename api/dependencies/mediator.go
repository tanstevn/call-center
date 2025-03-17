package dependencies

import (
	"call-center/internal/application/calls/commands"
	"call-center/internal/infrastructure/data/repositories"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func RegisterMediatorHandlers() error {
	callRepository := repositories.NewCallRepository()

	createCallCommandHandler := commands.NewCreateCallCommandHandler(callRepository)

	handlers := map[string]error{
		"CreateCallCommandHandler": mediatr.RegisterRequestHandler(createCallCommandHandler),
	}

	for name, err := range handlers {
		if err != nil {
			log.Fatalf("Failed to register %s: %v", name, err)
			return err
		}
	}

	log.Println("All mediator handlers are reigstered.")
	return nil
}
