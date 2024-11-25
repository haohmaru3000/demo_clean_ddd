package productmysql

import (
	"context"

	"demo_clean_ddd/module/product/domain"
)

func (repo MysqlRepository) CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error {
	// Single responsibility (Làm đúng nhiệm vụ chỉ có 1 dòng rồi return)
	if err := repo.db.Table(prod.TableName()).Create(&prod).Error; err != nil {
		return err
	}
	return nil
}
