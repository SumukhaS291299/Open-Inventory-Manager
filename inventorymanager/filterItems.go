package inventorymanager

// Add filter with go-routines in future

// import (
// 	"strings"
// )

// // --- Fetch logic ---
// func FindItems(filter ItemFilter) ([]*InventoryItem, error) {

// 	if c == nil || len(c.Items) == 0 {
// 		return nil, nil
// 	}

// 	var results []*InventoryItem

// 	for _, item := range c.Items {
// 		if item == nil || item.Attributes == nil {
// 			continue
// 		}

// 		attr := item.Attributes
// 		match := true

// 		// --- Check filters one by one ---
// 		if filter.ID != nil && *filter.ID != item.ID {
// 			match = false
// 		}
// 		if filter.Name != "" && !strings.EqualFold(attr.Name, filter.Name) {
// 			match = false
// 		}
// 		if filter.Category != "" && !strings.EqualFold(attr.Category, filter.Category) {
// 			match = false
// 		}
// 		if filter.Color != "" && !strings.EqualFold(attr.Color, filter.Color) {
// 			match = false
// 		}
// 		if filter.Location != "" && !strings.EqualFold(attr.Location, filter.Location) {
// 			match = false
// 		}
// 		if filter.IsActive != nil && attr.IsActive != *filter.IsActive {
// 			match = false
// 		}
// 		if filter.IsAvailable != nil && attr.IsAvailable != *filter.IsAvailable {
// 			match = false
// 		}
// 		if filter.Supplier != "" && item.Supplier != nil &&
// 			!strings.EqualFold(item.Supplier.Name, filter.Supplier) {
// 			match = false
// 		}

// 		if match {
// 			results = append(results, item)
// 		}
// 	}

// 	return results, nil
// }
