package inventorymanager

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/dgraph-io/badger"
)

var DB *badger.DB

func InitDB(path string, exit chan bool) error {
	var err error
	opts := badger.DefaultOptions(path)
	opts.Truncate = true

	DB, err = badger.Open(opts)
	if err != nil {
		log.Fatalf("Failed to open Badger DB: %v", err)
		close(exit) //Stop everything if something is wrong in persistence
		return err
	}

	return nil
}

func SaveToBadger(ID int64, data []byte) error {
	key := []byte(strconv.FormatInt(ID, 10))

	return DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, data)
	})
}

func LoadAllItems() error {
	return DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = true // fast for full scan
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()

			// Read value and unmarshal
			err := item.Value(func(val []byte) error {
				invItem := &InventoryItem{}
				if err := json.Unmarshal(val, invItem); err != nil {
					return err
				}

				// Add to in-memory Inv
				Inv.Items = append(Inv.Items, invItem)
				return nil
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}
