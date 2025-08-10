package inventrymanagerserver

import (
	"context"
	"log"
	"time"

	pb "openinventorymanager/openinventorymanager"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewInventoryManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create an item
	r, err := c.CreateItem(ctx, &pb.ItemInput{
		Name:        "Laptop",
		Location:    "Home Office",
		Description: "MacBook Pro",
		Tags:        []string{"electronics", "computer"},
		Metadata:    map[string]string{"brand": "Apple", "year": "2021"},
	})
	if err != nil {
		log.Fatalf("could not create item: %v", err)
	}
	log.Printf("Created Item: %v", r)

	// List items
	list, err := c.ListItems(ctx, &pb.ListItemsRequest{})
	if err != nil {
		log.Fatalf("could not list items: %v", err)
	}
	log.Printf("Inventory: %v", list.Items)
}
