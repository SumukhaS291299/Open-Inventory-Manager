package inventrymanagerserver

import (
	"context"
	"log"
	"net"

	pb "openinventorymanager/openinventorymanager"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type inventoryServer struct {
	pb.UnimplementedInventoryManagerServer
	items map[string]*pb.ItemOutput
}

func newServer() *inventoryServer {
	return &inventoryServer{items: make(map[string]*pb.ItemOutput)}
}

func (s *inventoryServer) CreateItem(ctx context.Context, in *pb.ItemInput) (*pb.ItemOutput, error) {
	id := "item123" // In real code, generate unique ID
	item := &pb.ItemOutput{
		Id:          id,
		Name:        in.Name,
		Location:    in.Location,
		Description: in.Description,
		Tags:        in.Tags,
		Metadata:    in.Metadata,
	}
	s.items[id] = item
	log.Printf("Created item: %+v", item)
	return item, nil
}

func (s *inventoryServer) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.ItemOutput, error) {
	item, exists := s.items[in.Id]
	if !exists {
		return nil, nil
	}
	return item, nil
}

func (s *inventoryServer) ListItems(ctx context.Context, in *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	list := &pb.ListItemsResponse{}
	for _, item := range s.items {
		list.Items = append(list.Items, item)
	}
	return list, nil
}

func RunServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryManagerServer(grpcServer, newServer())

	// Enable reflection so that it can discover services
	reflection.Register(grpcServer)

	log.Println("Server listening at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
