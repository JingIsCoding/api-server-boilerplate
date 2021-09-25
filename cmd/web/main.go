package main

import (
	"web-server/logger"
)

func main() {
	logger.Init()
	server, err := NewServer()
	if err != nil {
		logger.Errorf("Failed to start server. %v", err)
	}
	if err = server.Run(); err != nil {
		logger.Error("Failed to run server", err)
	}
}
