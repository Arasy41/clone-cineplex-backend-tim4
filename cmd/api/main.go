package main

import (
	"github.com/nsrvel/golang-example/config"
	"github.com/nsrvel/golang-example/internal/server"
	"github.com/nsrvel/golang-example/pkg/db"
)

func main() {

	//* Initial Configuration
	cfg := config.InitConfig()

	//* Connect Database
	database := db.NewDBConnection(&cfg.Database)

	server.RunFiber(cfg, database)
}
