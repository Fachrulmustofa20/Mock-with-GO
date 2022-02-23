package repository

import "mock/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
