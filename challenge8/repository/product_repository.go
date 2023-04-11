package repository

import (
	"challenge8/entity"
)

type ProductRepository interface {
	FindById(id string) *entity.Product
}
