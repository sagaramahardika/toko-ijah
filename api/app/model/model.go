package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Item struct {
	ID        int64  `gorm:"AUTO_INCREMENT"`
	SKU       string `gorm:"not null;unique" json:"sku"`
	Name      string `gorm:"not null" json:"name"`
	Quantity  int64  `gorm:"not null" json:"quantity"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
}

type IncomingItem struct {
	ID            int64  `gorm:"AUTO_INCREMENT"`
	SKU           string `gorm:"not null" json:"sku"`
	TotalOrdered  int64  `gorm:"not null" json:"totalordered"`
	TotalIncoming int64  `gorm:"not null" json:"totalincoming"`
	Price         int64  `gorm:"not null" json:"price"`
	TotalPrice    int64  `gorm:"not null" json:"totalprice"`
	ReceiptNumber string `gorm:"not null" json:"receiptnumber"`
	Note          string `json:"note,omitempty"`
	CreatedAt     int64  `gorm:"not null" json:"createdat"`
}

type OutgoingItem struct {
	ID            int64  `gorm:"AUTO_INCREMENT"`
	SKU           string `gorm:"not null" json:"sku"`
	TotalOutgoing int64  `gorm:"not null" json:"totaloutgoing"`
	Price         int64  `gorm:"not null" json:"price"`
	TotalPrice    int64  `gorm:"not null" json:"totalprice"`
	Note          string `json:"note,omitempty"`
	CreatedAt     int64  `gorm:"not null" json:"createdat"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&IncomingItem{})
	db.AutoMigrate(&OutgoingItem{})
	return db
}
