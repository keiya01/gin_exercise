package main

import (
	"github.com/keiya01/ginsampleapp/config"
	"github.com/keiya01/ginsampleapp/config/db"
	"github.com/keiya01/ginsampleapp/config/migration"

	_ "github.com/gin-gonic/gin"
)

func main() {
	setDB := db.DBInit()
	defer setDB.Close()
	migration.DBmigrate(setDB)
	config.Start()
}
