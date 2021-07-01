package service

import (
	"kraicklist/repository"
	"errors"
	"kraicklist/entity"

)

type CategoryService struct {
	Respository repository.CategoryRepository
}

func (service CategoryService) Get(id string)(*entity.Category, error)
	category :=service.Respository.FindById(id)
	if category == nil {
		return nil, errors.New("category not found")
	}else{
		return category,nil
	}
}