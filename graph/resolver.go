package graph

import (
	"github.com/emejotaw/graphql-api/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	productService      *service.ProductService
	orderService        *service.OrderService
	orderProductService *service.OrderProductService
}

func NewResolver(
	productService *service.ProductService,
	orderService *service.OrderService,
	orderProductService *service.OrderProductService,
) *Resolver {
	return &Resolver{
		productService:      productService,
		orderService:        orderService,
		orderProductService: orderProductService,
	}
}
