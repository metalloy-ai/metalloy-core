package main

import (
	"logiflowCore/internal/api"
	"logiflowCore/internal/config"
	"logiflowCore/pkg/validator"
)

func main() {
	cfg := config.LoadBaseConfig()
	validator.ValidateConfig()
	server := api.InitServer(*cfg)
	server.LoadServerConfig()
	server.Run()
}
