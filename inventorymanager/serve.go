package inventorymanager

import (
	"context"
	"log"
	"net/http"
	"openinventorymanager/logger"
	"time"

	"github.com/gin-gonic/gin"
)

var ginengine *gin.Engine

func init() {
	ginengine = gin.Default()

}

func getItems(c *gin.Context) {
	category := c.Query("category")

	c.JSON(200, struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}{
		Name:     "All items",
		Category: category,
	})
}

func EnableServices() {
	ginengine.GET("/items", getItems)
}

func Run(exit <-chan bool) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: ginengine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-exit // blocking

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
	logger.Logger.Info("Server was gracefully shutdown")
}
