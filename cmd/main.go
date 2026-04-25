package main

import (
	"tipe-handling/config"
	boot "tipe-handling/pkg/bootstrap"
)

func main() {
	config.LoadEnv()

	app := boot.NewApp()
	app.Run()
}