package main

import "logiflowCore/internal/api"

func main() {
	server := api.InitServer()
	server.LoadConfig()
	server.Run()
}