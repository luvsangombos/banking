package main

import (
	"github.com/luvsangombos/banking/app"
	"github.com/luvsangombos/banking/logger"
)

func main() {

	logger.Info("Starting the application")
	setenv()
	app.Start()
}
