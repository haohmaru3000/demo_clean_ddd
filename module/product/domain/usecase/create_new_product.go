package productusecase

import (
	"context"
	"strings"

	"demo_clean_ddd/module/product/domain"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod productdomain.ProductCreationDTO) error
}

func NewCreateProductUseCase(repo CreateProductRepository) CreateNewProductUseCase {
	return CreateNewProductUseCase{
		repo: repo,
	}
}

type CreateNewProductUseCase struct {
	repo CreateProductRepository
}

// Interface ở đâu thì dùng hàm của nó ở đó
func (uc CreateNewProductUseCase) CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error {
	/** Business Logic: do team Product ràng buộc (nên code trc tiên) **/
	// - Phần nhiều nhất mà ta cần maintain, đảm bảo tính chính xác cho Unit-test ở phần này
	// EX: check tính đúng đắn dữ liệu, hoặc gồm cả qui trình(bao nhiêu steps khi tạo đơn hàng, chuyển khoản...)
	prod.Name = strings.TrimSpace(prod.Name)

	if prod.Name == "" {
		return productdomain.ErrProductNameCannotBeBlank
	}

	if err := uc.repo.CreateProduct(ctx, prod); err != nil {
		return err
	}

	return nil
}

type CreateProductRepository interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error
}
