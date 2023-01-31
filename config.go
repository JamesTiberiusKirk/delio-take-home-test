package main

import (
	"log"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ApiKey string `env:"API_KEY"`
}

func buildConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c Config
	_, err = env.UnmarshalFromEnviron(&c)
	if err != nil {
		logrus.
			WithError(err).
			Fatal("unable to get env variables")
	}

	return c
}
