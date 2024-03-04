package cmd

import (
	"log"

	"binod210/task_remainder/config"
	"binod210/task_remainder/db"
	"binod210/task_remainder/https"
)

// read config and start server
func Execute() {
	config.LoadDotEnv()
	configs := config.GetConfig()

	port := configs.App.Port
	isDebug := configs.App.Debug
	_, err := db.ConnectDB(configs.Database.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.Migrate()

	https.StartServer(port, isDebug)
}
