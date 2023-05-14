package config

import (
	"os"
	"runtime"
	"strconv"
)

type Setting struct {
	Host          string
	Port          int
	ApiVersion    string
	Version       string
	PG_URL        string
	REDIS_URL     string
	REDIS_PWS     string
	Env           string
	NumCPU        int
	TokenDuration int
	JwtKey        string
}

func LoadBaseConfig() *Setting {
	LoadEnv(".env")

	version := os.Getenv("VERSION")
	apiVersion := os.Getenv("API_VERSION")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	PG_URL := os.Getenv("PG_URL")
	REDIS_URL := os.Getenv("REDIS_URL")
	REDIS_PWS := os.Getenv("REDIS_PSW")
	env := os.Getenv("ENV")
	numCPU := runtime.NumCPU()
	tokenDuration, _ := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	jwtKey := os.Getenv("JWT_KEY")

	return &Setting{
		Host:       host,
		Port:       port,
		ApiVersion: apiVersion,
		Version:    version,
		PG_URL:     PG_URL,
		REDIS_URL:  REDIS_URL,
		REDIS_PWS:  REDIS_PWS,
		Env:        env,
		NumCPU:     numCPU,
		TokenDuration: tokenDuration,
		JwtKey: jwtKey,
	}
}
