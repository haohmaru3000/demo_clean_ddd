package controller

import (
	"gorm.io/gorm"
	"net/http"

	"demo_clean_ddd/common"
	productdomain "demo_clean_ddd/module/product/domain"

	"github.com/gin-gonic/gin"
)

// Có thể viết theo encapsulation hoặc kiểu method
func (api APIController) CreateProductAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check & parse data from body
		var productData productdomain.ProductCreationDTO

		if err := c.Bind(&productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		productData.Id = common.GenUUID()

		if err := api.createUseCase.CreateProduct(c.Request.Context(), &productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// response to client
		c.JSON(http.StatusCreated, gin.H{"data": productData.Id})
	}
}
