package main

import (
	"EM-TASK/app/internal/app"
	"flag"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	"log"
)

// @title EffectiveMobileTask
// @host localhost:3005
func main() {
	readEnv := flag.String("env-file", ".env", "Path to application env file.")
	configPath := flag.String("config-path", "config.yml", "Path to application config file.")
	flag.Parse()
	println(*configPath)
	if *readEnv != "" {
		if err := godotenv.Load(*readEnv); err != nil {
			log.Fatal(err)
		}
	}

	application, err := app.New(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err = application.Start(); err != nil {
		log.Fatal(err)
	}
}
