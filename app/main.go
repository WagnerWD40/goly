package main

import (
	"app/model"
	"app/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
