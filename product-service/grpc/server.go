package grpc

import (
	"context"

	"github.com/prashant-bhilwar/order-processing-system/product-service/gen"
	"github.com/prashant-bhilwar/order-processing-system/product-service/repository"
)

type ProductGRPCServer struct {
	gen.UnimplementedProductServiceServer
}

func (s *ProductGRPCServer) GetAllProducts(ctx context.Context, _ *gen.Empty) (*gen.ProductList, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	var grpcProducts []*gen.Product
	for _, p := range products {
		grpcProducts = append(grpcProducts, &gen.Product{
			Id:          int32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Price:       float32(p.Price),
			InStock:     p.InStock,
		})
	}
	return &gen.ProductList{Products: grpcProducts}, nil
}

func (s *ProductGRPCServer) GetProductById(ctx context.Context, req *gen.ProductId) (*gen.Product, error) {
	product, err := repository.GetProductbyID(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &gen.Product{
		Id:    int32(product.ID),
		Name:  product.Name,
		Price: float32(product.Price),
	}, nil
}
