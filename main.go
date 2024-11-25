package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	"demo_clean_ddd/module/product/controller"
	"demo_clean_ddd/util"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	config, err := util.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("cannot load config:", err)
// 	}
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		config.DBUserName,
// 		config.DBUserPassword,
// 		config.DBHost,
// 		config.DBPort,
// 		config.DBName,
// 	)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	db = db.Debug()

// 	now := time.Now().UTC() // Will use GMT+7 if no UTC()

// 	newId, _ := uuid.NewV7()

// 	newProd := Product{
// 		BaseModel: BaseModel{
// 			Id:        newId,
// 			Status:    "activated",
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		CategoryId:  1,
// 		Name:        "Latte",
// 		Image:       nil,
// 		Type:        "drink",
// 		Description: "",
// 	}

// 	if err := db.Table(Product{}.TableName()).Create(&newProd).Error; err != nil {
// 		log.Println(err)
// 	}

// 	var oldProduct Product

// 	if err := db.
// 		Table(Product{}.TableName()).
// 		// Where("id = ?", 3).
// 		First(&oldProduct).Error; err != nil {
// 		log.Println(err)
// 	}

// 	log.Println("Product: ", oldProduct)

// 	var prods []Product

// 	if err := db.
// 		Table(Product{}.TableName()).
// 		Where("status not in (?)", []string{"deactivated"}).
// 		Limit(10).
// 		Offset(10).
// 		Order("id desc").
// 		Find(&prods).Error; err != nil {
// 		log.Println(err)
// 	}

// 	log.Println("Products: ", prods)

// 	// Update 'oldProduct' using map (must provide "Where condition")
// 	if err := db.
// 		Table(Product{}.TableName()).
// 		Where("id = ?", 3).
// 		Updates(map[string]any{"name": "Cappuccino"}).Error; err != nil {
// 		log.Println(err)
// 	}

// 	// oldProduct.Name = "" // Update its name to ""

// 	// Default update (no need to use "Where condition" cuz oldProduct has provided its 'Id')
// 	// if err := db.
// 	// 	Table(Product{}.TableName()).
// 	// 	Updates(oldProduct).Error; err != nil {
// 	// 	log.Println(err) // Will cause error cuz Gorm won't update "empty, nil, 0" values
// 	// }

// 	// emptyStr := ""

// 	// if err := db.
// 	// 	Table(Product{}.TableName()).
// 	// 	Where("id = ?", 3).
// 	// 	Updates(ProductUpdate{Name: &emptyStr}).Error; err != nil {
// 	// 	log.Println(err)
// 	// }

// 	// Delete
// 	// if err := db.
// 	// 	Table(Product{}.TableName()).
// 	// 	Where("id = ?", 4).
// 	// 	Delete(nil).Error; err != nil {
// 	// 	log.Println(err)
// 	// }
// }

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUserName,
		config.DBUserPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", controller.CreateProductAPI(db))
		}
	}

	r.Run(":3000")
}
