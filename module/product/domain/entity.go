package productdomain

import "demo_clean_ddd/common"

type Product struct {
	common.BaseModel
	CategoryId int    `gorm:"column:category_id" json:"category_id"`
	Name       string `gorm:"column:name" json:"name"`
	//Image       any    `gorm:"column:image"`
	Type        string `gorm:"column:type" json:"type"`
	Description string `gorm:"column:description" json:"description"`
}
