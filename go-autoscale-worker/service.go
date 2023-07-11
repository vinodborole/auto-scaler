package main

import (
	"fmt"
	"log"
	"strings"

	pb "github.com/vinodborole/go-autoscale-worker/proto"
)

// service is used to implement ProductInfo.
type productService struct {
	pb.UnimplementedProductInfoServer
	productMap map[string]*pb.Product
}

// SearchOrders implements ProductInfo.SearchProducts
func (s *productService) SearchOrders(searchQuery *pb.SearchQuery, stream pb.ProductInfo_SearchProductsServer) error {
	for key, product := range s.productMap {
		log.Print(key, product)

		if strings.Contains(product.Description, searchQuery.Value) {
			// Send the matching orders in a stream
			err := stream.Send(product)
			if err != nil {
				return fmt.Errorf("error sending message to stream: %v", err)
			}
			log.Print("Matching Product Found : " + key)
			break
		}

	}
	return nil
}
