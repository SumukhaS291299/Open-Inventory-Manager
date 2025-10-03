package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		fmt.Printf("A error occurred: %s", err)
	}
	defer Logger.Sync() // flushes buffer, if any
}
