package models

import "time"

type Client struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Cuit      string    `gorm:"column:cuit;unique;not null" json:"cuit"`
	Contact   string    `json:"contact"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Orders    []Order   `json:"orders,omitempty" gorm:"foreignKey:ClientID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
