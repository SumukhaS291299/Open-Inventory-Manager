package main

import (
	"openinventorymanager/inventorymanager"
	"openinventorymanager/logger"
	"time"
)

func main() {

	exit := make(chan bool) // From any other operations

	logger.InitLogger()
	logger.Logger.Info("Starting application...")
	// qrutils.Generate("www.google.com")
	inventorymanager.EnableServices()
	go inventorymanager.Run(exit)

	time.Sleep(time.Second * 5)
	<-exit
	logger.Logger.Info("Waiting all process to stop gracefully")
}
