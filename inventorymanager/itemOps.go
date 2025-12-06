package inventorymanager

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/SumukhaS291299/Open-Inventory-Manager/logger"
)

var (
	items         *InventoryCollection
	counter       int64 = 0
	lastTimestamp int64 = 0
)

// --- ID generator ---
func generateUniqueID() int64 {
	// Only unique IDs are using UTC time
	// All others are using Local time
	now := time.Now().UnixMilli()
	if now == atomic.LoadInt64(&lastTimestamp) {
		c := atomic.AddInt64(&counter, 1)
		return now*1000 + c
	} else {
		atomic.StoreInt64(&lastTimestamp, now)
		atomic.StoreInt64(&counter, 0)
		return now * 1000
	}
}

// --- Add Item ---
func (c *InventoryCollection) AddItem(unitPrice float32, stock int, name, description, location, color, category string) (*InventoryItem, error) {
	if name == "" || category == "" {
		return nil, errors.New("name and category are required")
	}

	item := &InventoryItem{
		ID: generateUniqueID(),
		Attributes: &Attributes{
			Name:        name,
			Description: description,
			Color:       color,
			Category:    category,
			UnitPrice:   unitPrice,
			StockLevel:  stock,
			Location:    location,
			IsActive:    true,
			IsAvailable: stock > 0,
		},
		TimeMeta: &TimeMeta{
			Bought:   time.Now().Local(),
			Modified: time.Now().Local(),
		},
	}

	c.mu.Lock()
	c.Items = append(c.Items, item)
	c.mu.Unlock()

	go func(id int64, localItem *InventoryItem) {
		data, err := json.Marshal(localItem)
		if err != nil {
			logger.Logger.Error("Failed to serialize item: " + err.Error())
		}
		saveErr := SaveToBadger(id, data)
		if saveErr != nil {
			logger.Logger.Error(saveErr.Error())
		} else {
			logger.Logger.Debug("Saved item with ID: \t" + strconv.FormatInt(id, 10) + "\n\t With name:" + item.Attributes.Name)
		}

	}(item.ID, item)

	return item, nil
}

// --- Find / Filter Items ---
func (c *InventoryCollection) FindItems(filter ItemFilter) []*InventoryItem {
	c.mu.Lock()
	defer c.mu.Unlock()

	var results []*InventoryItem
	for _, item := range c.Items {
		if item == nil || item.Attributes == nil {
			continue
		}

		attr := item.Attributes
		match := true

		if filter.ID != nil && *filter.ID != item.ID {
			match = false
		}
		if filter.Name != "" && !strings.EqualFold(attr.Name, filter.Name) {
			match = false
		}
		if filter.Category != "" && !strings.EqualFold(attr.Category, filter.Category) {
			match = false
		}
		if filter.Color != "" && !strings.EqualFold(attr.Color, filter.Color) {
			match = false
		}
		if filter.Location != "" && !strings.EqualFold(attr.Location, filter.Location) {
			match = false
		}
		if filter.IsActive != nil && attr.IsActive != *filter.IsActive {
			match = false
		}
		if filter.IsAvailable != nil && attr.IsAvailable != *filter.IsAvailable {
			match = false
		}
		if filter.Supplier != "" && item.Supplier != nil &&
			!strings.EqualFold(item.Supplier.Name, filter.Supplier) {
			match = false
		}

		if match {
			results = append(results, item)
		}
	}

	return results
}

// --- Modify Item (by ID) ---
func (c *InventoryCollection) ModifyItem(id int64, updated *Attributes) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, item := range c.Items {
		if item.ID == id {
			if updated.Name != "" {
				item.Attributes.Name = updated.Name
			}
			if updated.Description != "" {
				item.Attributes.Description = updated.Description
			}
			if updated.Color != "" {
				item.Attributes.Color = updated.Color
			}
			if updated.Category != "" {
				item.Attributes.Category = updated.Category
			}
			if updated.Location != "" {
				item.Attributes.Location = updated.Location
			}
			if updated.UnitPrice != 0 {
				item.Attributes.UnitPrice = updated.UnitPrice
			}
			if updated.StockLevel >= 0 {
				item.Attributes.StockLevel = updated.StockLevel
				item.Attributes.IsAvailable = updated.StockLevel > 0
			}

			item.TimeMeta.Modified = time.Now().Local()
			return nil
		}
	}
	return errors.New("item not found")
}

// --- Delete Item (by ID) ---
func (c *InventoryCollection) DeleteItem(id int64) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, item := range c.Items {
		if item.ID == id {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("item not found")
}

// --- Get Item IDs based on filters ---
func (c *InventoryCollection) getItemIDs(filter ItemFilter) []int64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	var ids []int64

	for _, item := range c.Items {
		if item == nil || item.Attributes == nil {
			continue
		}

		attr := item.Attributes
		match := true

		if filter.ID != nil && *filter.ID != item.ID {
			match = false
		}
		if filter.Name != "" && !strings.EqualFold(attr.Name, filter.Name) {
			match = false
		}
		if filter.Category != "" && !strings.EqualFold(attr.Category, filter.Category) {
			match = false
		}
		if filter.Color != "" && !strings.EqualFold(attr.Color, filter.Color) {
			match = false
		}
		if filter.Location != "" && !strings.EqualFold(attr.Location, filter.Location) {
			match = false
		}
		if filter.IsActive != nil && attr.IsActive != *filter.IsActive {
			match = false
		}
		if filter.IsAvailable != nil && attr.IsAvailable != *filter.IsAvailable {
			match = false
		}
		if filter.Supplier != "" && item.Supplier != nil &&
			!strings.EqualFold(item.Supplier.Name, filter.Supplier) {
			match = false
		}

		if match {
			ids = append(ids, item.ID)
		}
	}

	return ids
}
