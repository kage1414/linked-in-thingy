package main

import (
	"log"

	"job-board/backend/server"
)

func main() {
	s := server.NewServer()
	if err := s.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
