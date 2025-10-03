package inventorymanager

import (
	"sync/atomic"
	"time"
)

var items *InventoryCollection

var counter int64 = 0
var lastTimestamp int64 = 0

func generateUniqueID() int64 {
	now := time.Now().UnixMilli()
	if now == atomic.LoadInt64(&lastTimestamp) {
		// same millisecond → increment counter
		c := atomic.AddInt64(&counter, 1)
		return now*1000 + c // shift counter to keep uniqueness
	} else {
		atomic.StoreInt64(&lastTimestamp, now)
		atomic.StoreInt64(&counter, 0)
		return now * 1000
	}
}

func AddAttribute(name, description, color, category, location, photoBase64 string, unitPrice float32, stockLevel int, isActive, isAvailable bool) Attributes {
	return Attributes{Name: name, Description: description, Color: color, Category: category, Location: location, PhotoBase64: photoBase64, UnitPrice: unitPrice, StockLevel: stockLevel, IsActive: isActive, IsAvailable: isAvailable}
}

// All date and time in local time
func AddTimeMeta(bought, expires time.Time) TimeMeta {
	return TimeMeta{Modified: time.Now().Local(), Bought: bought, Expires: expires}

}

// func AddItemMeta(Operation, ResourceType, RequestID string) TimeMeta {

// }

func AddSupplier(supplierType, name, address string, Online bool) Supplier {

	return Supplier{SupplierType: supplierType, Name: name, Address: address}

}

func AddItem(UnitPrice float32, stock int, name, description, location, color, Category string) (added bool, err error) {

	item := InventoryItem{ID: generateUniqueID()}

	items.Items = append(items.Items, &item)

	return false, nil

}

// TODO Add cache when new items are added as append
