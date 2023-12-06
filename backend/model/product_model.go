package model

type Product struct {
	Id          int     `gorm:"type:int(11);"`
	Title       string  `gorm:"type:varchar(2500);"`
	Image       string  `gorm:"type:varchar(250);"`
	Price       float32 `gorm:"type:float;"`
	Description string  `gorm:"type:varchar(2500);"`
}

type Products []Product
