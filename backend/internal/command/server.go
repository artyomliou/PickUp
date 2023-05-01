package command

import (
	"fmt"
	_ "the-video-project/backend/configs"
	"the-video-project/backend/internal/httpapi"
)

type ServerCommand struct {
	Port int
}

func (cmd ServerCommand) Run() error {
	router := httpapi.SetupRouter()
	return router.Run(fmt.Sprintf(":%d", cmd.Port))
}
