package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	Id        uuid.UUID `gorm:"column:id;" json:"id"`
	Status    string    `gorm:"column:status;" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`
}

func GenNewModel() BaseModel {
	now := time.Now().UTC() // Will use GMT+7 if no UTC()
	newId, _ := uuid.NewV7()

	return BaseModel{
		Id:        newId,
		Status:    "activated",
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func GenUUID() uuid.UUID {
	newId, _ := uuid.NewV7()
	return newId
}
