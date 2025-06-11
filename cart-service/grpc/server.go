package grpc

import (
	"context"

	"github.com/prashant-bhilwar/order-processing-system/cart-service/gen"
	"github.com/prashant-bhilwar/order-processing-system/cart-service/repository"
)

type CartGRPCServer struct {
	gen.UnimplementedCartServiceServer
}

func (s *CartGRPCServer) GetCartItems(ctx context.Context, req *gen.UserRequest) (*gen.CartItemList, error) {
	items, err := repository.GetCartItemsByUserID(uint(req.UserId))
	if err != nil {
		return nil, err
	}

	var grpcItems []*gen.CartItem
	for _, item := range items {
		grpcItems = append(grpcItems, &gen.CartItem{
			Id:        int32(item.ID),
			UserId:    int32(item.UserID),
			ProductId: int32(item.ProductID),
			Quantity:  int32(item.Quantity),
		})
	}
	return &gen.CartItemList{Items: grpcItems}, nil
}
