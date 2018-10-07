package main

import (
	"toko-ijah/api/app"
	"toko-ijah/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
