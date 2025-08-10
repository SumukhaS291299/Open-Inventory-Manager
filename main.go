package main

import (
	inventrymanagerserver "openinventorymanager/inventryManagerServer"
	"openinventorymanager/logger"
)

func main() {
	logger.InitLogger()
	logger.Logger.Info("Starting application...")
	// qrutils.Generate("www.google.com")
	inventrymanagerserver.RunServer()
}
