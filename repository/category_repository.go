package repository

import "kraicklist/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
