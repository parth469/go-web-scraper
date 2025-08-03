package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Url  string `json:"url"`
	Port string `json:"port"`
}

var Env Environment

func Init() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	Env = Environment{
		Url:  os.Getenv("URL"),
		Port: os.Getenv("PORT"),
	}


	if Env.Port == "" || Env.Url == "" {
		return fmt.Errorf("environment variables missing: URL or PORT is not set")
	}

	return nil
}
