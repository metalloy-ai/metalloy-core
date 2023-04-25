package main

import (
	"metalloyCore/internal/api"
	"metalloyCore/internal/config"
	"metalloyCore/pkg/validator"
)

func main() {
	cfg := config.LoadBaseConfig()
	validator.ValidateConfig()
	server := api.InitServer(*cfg)
	server.LoadServerConfig()
	server.Run()
}
