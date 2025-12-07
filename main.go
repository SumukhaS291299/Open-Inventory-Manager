package main

import (
	"flag"
	"time"

	"github.com/SumukhaS291299/Open-Inventory-Manager/inventorymanager"
	"github.com/SumukhaS291299/Open-Inventory-Manager/logger"
)

func main() {

	exit := make(chan bool) // From any other operations

	dbPath := flag.String("dbpath", "", "Path to DB folder")
	flag.Parse()

	if *dbPath == "" {
		logger.Logger.Error("Error: --dbpath is required")
	}
	err := inventorymanager.InitDB(*dbPath, exit)

	if err != nil {
		logger.Logger.Error(err.Error())
		logger.Logger.Fatal("There was some error in loading your data")
		close(exit)
	}

	err = inventorymanager.LoadAllItems()

	if err != nil {
		logger.Logger.Error(err.Error())
		logger.Logger.Fatal("There was some error in loading your data")
		close(exit)
	}

	logger.InitLogger()
	logger.Logger.Info("Starting application...")

	// qrutils.Generate("www.google.com")
	inventorymanager.EnableServices()
	go inventorymanager.Run(exit)

	val, ok := <-exit
	if ok && val {
		close(exit) // optional: close channel if you want to signal others
	}
	// Sleep for a minute so that all go routines can come to stop
	time.Sleep(time.Minute)
}
