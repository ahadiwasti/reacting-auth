package main

import (
	"github.com/ahadiwasti/reacting-auth/cmd"
)

// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
func main() {
	cmd.Execute()
	//dao.Shutdown()
}
