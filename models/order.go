package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	OrderNumber  string    `gorm:"unique" json:"order_number"`
	ClientID     uint      `gorm:"not null" json:"client_id"`
	Client       Client    `json:"client,omitempty" gorm:"foreignKey:ClientID"`
	ProductType  string    `json:"product_type"`
	Material     string    `json:"material"`
	ThicknessMM  float64   `json:"thickness_mm"`
	WidthMM      float64   `json:"width_mm"`
	LengthMM     float64   `json:"length_mm"`
	WeightKG     float64   `json:"weight_kg"`
	Quantity     int       `json:"quantity"`
	OrderDate    time.Time `json:"order_date"`
	DeliveryDate time.Time `json:"delivery_date"`
	Status       string    `json:"status"`
	Notes        string    `json:"notes"`
	UnitPrice    float64   `json:"unit_price"`
	TotalPrice   float64   `json:"total_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
