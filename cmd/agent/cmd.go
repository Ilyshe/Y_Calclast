package main

import (
	"github.com/Ilyshe/Y_Calclast/internal/agent"
	"github.com/Ilyshe/Y_Calclast/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	agent := agent.New(cfg)
	agent.Run()
}
