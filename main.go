package main

import (
	"lolyshop/config"
	"lolyshop/modules/servers"
	"lolyshop/pkg/databases"
	"os"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env.dev"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
