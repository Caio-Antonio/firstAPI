package main

import (
	"log/slog"

	"github.com/Caio-Antonio/firstAPI/config/logger"
)

func main() {
	logger.InitLogger()
	slog.Info("starting api")
}