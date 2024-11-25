package controller

import (
	"gorm.io/gorm"
	"net/http"

	productdomain "demo_clean_ddd/module/product/domain"
	productusecase "demo_clean_ddd/module/product/domain/usecase"
	productmysql "demo_clean_ddd/module/product/repository/mysql"

	"github.com/gin-gonic/gin"
)

// Có thể viết theo encapsulation hoặc kiểu method
func CreateProductAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check & parse data from body
		var productData productdomain.ProductCreationDTO

		if err := c.Bind(&productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		repo := productmysql.NewMysqlRepository(db)
		useCase := productusecase.NewCreateProductUseCase(repo)

		if err := useCase.CreateProduct(c.Request.Context(), &productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// response to client
		c.JSON(http.StatusCreated, gin.H{"data": productData.Id})
	}
}
