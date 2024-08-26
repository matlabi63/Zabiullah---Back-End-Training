package models

import (
	"github.com/google/uuid"
)

type NutritionalData struct {
	Calories      float64 `json:"calories"`
	Fat           float64 `json:"fat"`
	Sugar         float64 `json:"sugar"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
}

type Fruit struct {
	ID             uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name           string          `json:"name"`
	Price          float64         `json:"price"`
	NutritionalData NutritionalData `json:"nutritional_data"`
}

func (Fruit) TableName() string {
	return "fruits"
}
