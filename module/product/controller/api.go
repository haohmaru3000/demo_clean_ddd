package controller

import (
	"context"

	productdomain "demo_clean_ddd/module/product/domain"
)

// Nhận vào tất cả các Use-Cases của nó
type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error
}

type APIController struct {
	createUseCase CreateProductUseCase
}

func NewAPIController(createUseCase CreateProductUseCase) APIController {
	return APIController{createUseCase: createUseCase}
}
