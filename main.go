package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/Andrey_Lebedinskiy/test_bot/config"
	"gitlab.com/Andrey_Lebedinskiy/test_bot/internal/infrastructure/logs"
	"gitlab.com/Andrey_Lebedinskiy/test_bot/run"
	"os"
)

func main() {
	err := godotenv.Load()
	conf := config.NewAppConf()
	logger := logs.NewLogger(conf, os.Stdout)
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	conf.Init(logger)

	app := run.NewApp(conf, logger)
	exitCode := app.Bootstrap().Run()
	os.Exit(exitCode)
}
