package main

import (
	"homework/internal/commander"
	"homework/internal/handlers"
	"log"
)

func main() {
	log.Println("start main")
	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}
	handlers.AddHandlers(cmd)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
