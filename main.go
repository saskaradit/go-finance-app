package main

import (
	"github.com/saskaradit/finance-app/app"
	"github.com/saskaradit/finance-app/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
