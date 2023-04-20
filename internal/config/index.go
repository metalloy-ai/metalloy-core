package config

import (
	"os"
	"strconv"
)

type Setting struct {
	Host       string
	Port       int
	ApiVersion string
	Version	   string
	PG_URL	   string
	MG_URL	   string
	Env        string
}

func LoadBaseConfig() *Setting {
	LoadEnv(".env")

	version := os.Getenv("VERSION")
	apiVersion := os.Getenv("API_VERSION")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	PG_URL := os.Getenv("PG_URL")
	MG_URL := os.Getenv("MG_URL")
	env := os.Getenv("ENV")

	return &Setting{
		Host:       host,
		Port:       port,
		ApiVersion: apiVersion,
		Version:    version,
		PG_URL:     PG_URL,
		MG_URL:     MG_URL,
		Env:        env,
	}
}
