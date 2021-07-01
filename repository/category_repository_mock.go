package repository

import (
	"kraicklist/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.mock
}

func (repository *CategoryRepository) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(Entity.Category)
		return &category
	}
}
