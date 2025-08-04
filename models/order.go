package models

import "time"

type Order struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	OrderNumber   string    `gorm:"unique" json:"order_number"`
	ClientName    string    `json:"client_name"`
	ClientCUIT    string    `gorm:"column:client_cuit" json:"client_cuit"`
	ClientContact string    `json:"client_contact"`
	ClientAddress string    `json:"client_address"`
	ProductType   string    `json:"product_type"`
	Material      string    `json:"material"`
	ThicknessMM   float64   `json:"thickness_mm"`
	WidthMM       float64   `json:"width_mm"`
	LengthMM      float64   `json:"length_mm"`
	WeightKG      float64   `json:"weight_kg"`
	Quantity      int       `json:"quantity"`
	OrderDate     time.Time `json:"order_date"`
	DeliveryDate  time.Time `json:"delivery_date"`
	Status        string    `json:"status"`
	Notes         string    `json:"notes"`
	UnitPrice     float64   `json:"unit_price"`
	TotalPrice    float64   `json:"total_price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
