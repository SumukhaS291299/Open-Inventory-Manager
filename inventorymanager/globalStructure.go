package inventorymanager

import (
	"sync"
	"time"
)

// Will add later is necessary
//  Meta       *ItemMeta   `json:"meta"`               // Metadata about operations
// type ItemMeta struct {
// 	Operation    string `json:"operation,omitempty"`
// 	ResourceType string `json:"resource_type,omitempty"`
// 	RequestID    string `json:"request_id"`
// }

// --- Core Item ---
type InventoryItem struct {
	ID         int64       `json:"id"`                 // Unique system-wide ID
	Attributes *Attributes `json:"attributes"`         // Core product details
	TimeMeta   *TimeMeta   `json:"time_meta"`          // Purchase + expiry
	Supplier   *Supplier   `json:"supplier,omitempty"` // Supplier
	Comments   []*Comment  `json:"comments,omitempty"` // Multiple, optional
	Tags       []*Tag      `json:"tags,omitempty"`     // Tags to filter custom definitions
}

// --- Attributes & Metadata ---
type Attributes struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Color       string  `json:"color,omitempty"`
	Category    string  `json:"category"`
	UnitPrice   float32 `json:"unit_price"`
	StockLevel  int     `json:"stock_level"`
	Location    string  `json:"location"`
	IsActive    bool    `json:"is_active,omitempty"`
	IsAvailable bool    `json:"is_available,omitempty"`
	PhotoBase64 string  `json:"photo_base64,omitempty"`
	Tags        []*Tag  `json:"tags,omitempty"`
}

// All date and time in local time
type TimeMeta struct {
	Bought   time.Time `json:"bought"`
	Expires  time.Time `json:"expires,omitempty"`
	Modified time.Time `json:"modified"`
}

// --- Linked Entities ---
type Supplier struct {
	SupplierType string `json:"supplier_type,omitempty"`
	Name         string `json:"name"`
	Online       bool   `json:"online,omitempty"`
	Address      string `json:"address,omitempty"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// --- Comments ---
type Comment struct {
	ID            string      `json:"id"`
	Miscellaneous interface{} `json:"miscellaneous"`
	CreatedAt     time.Time   `json:"created_at"`
	CreatedBy     string      `json:"created_by"`
}

// --- Collection with lock ---
type InventoryCollection struct {
	mu    sync.Mutex
	Items []*InventoryItem `json:"items"`
}

// --- Filter Definition ---
type ItemFilter struct {
	ID          *int64
	Name        string
	Category    string
	Color       string
	Location    string
	Supplier    string
	IsActive    *bool
	IsAvailable *bool
}

type ModifyRequest struct {
	Filter ItemFilter    `json:"filter"`
	Update InventoryItem `json:"update"`
}
