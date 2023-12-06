package dto

type ProductDto struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}

type ProductsDto []ProductDto
