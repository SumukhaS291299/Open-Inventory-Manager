package inventorymanager

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SumukhaS291299/Open-Inventory-Manager/logger"

	"github.com/gin-gonic/gin"
)

var ginengine *gin.Engine
var Inv *InventoryCollection

func init() {
	Inv = &InventoryCollection{}
	ginengine = gin.Default()

}

func addItem(c *gin.Context) {
	var item InventoryItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := Inv.AddItem(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	full := c.DefaultQuery("full", "false") == "true"
	if full {
		c.JSON(http.StatusOK, newItem)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "item added", "id": newItem.ID})
	}
}

func filterItems(c *gin.Context) {
	var filter ItemFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := Inv.FindItems(filter)
	c.JSON(http.StatusOK, results)
}

func getAllItems(c *gin.Context) {
	results := Inv.FindItems(ItemFilter{})
	c.JSON(http.StatusOK, results)
}

func modifyItem(c *gin.Context) {
	var req ModifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find matching items
	matches := Inv.FindItems(req.Filter)
	if len(matches) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no items match the filter"})
		return
	}
	if len(matches) > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filter matches multiple items; modify one at a time"})
		return
	}

	target := matches[0]

	// Apply updates
	if req.Update.Attributes != nil {
		_ = Inv.ModifyItem(target.ID, req.Update.Attributes)
	}

	if req.Update.Supplier != nil {
		target.Supplier = req.Update.Supplier
	}

	full := c.DefaultQuery("full", "false") == "true"
	if full {
		c.JSON(http.StatusOK, target)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "item modified"})
	}
}

func deleteItems(c *gin.Context) {
	var item InventoryItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if item.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	ids := Inv.getItemIDs(ItemFilter{ID: &item.ID})
	if len(ids) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	deleted := Inv.FindItems(ItemFilter{ID: &item.ID})[0]
	success, err := Inv.DeleteItem(item.ID)
	if err != nil || !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete item"})
		return
	}

	full := c.DefaultQuery("full", "false") == "true"
	if full {
		c.JSON(http.StatusOK, deleted)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
	}
}

// --- Enable Routes ---
func EnableServices() {
	ginengine.POST("/additem", addItem)          // create
	ginengine.DELETE("/deleteitem", deleteItems) // delete
	ginengine.POST("/filteritem", filterItems)   // read list
	ginengine.PUT("/modifyitem", modifyItem)     // update
	ginengine.GET("/getitems", getAllItems)      // get all items
}

func Run(exit chan bool) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: ginengine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
			exit <- false
		}
	}()

	<-exit // blocking here
	//  make sure the main thread is alive for few seconds after exit chan get's closed

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
	logger.Logger.Info("Server was gracefully shutdown")
}
