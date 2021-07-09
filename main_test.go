package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

type ItemRepositoryInterface interface {
	Search() ([]Record, error)
	SearchByCat() ([]Record, error)
}

type ItemService struct {
	ItemRepositoryInterface
}

func (mock *MockRepository) Search() ([]Record, error) {

	args := mock.Called()

	record := []Record{
		{ID: 12, Title: "Apple", Content: "Apple Handphone"},
	}

	return record, args.Error(1)
}

func (mock *MockRepository) SearchByCat() ([]Record, error) {

	args := mock.Called()

	record := []Record{
		{ID: 12, Title: "Apple", Content: "Apple Handphone"},
	}

	return record, args.Error(1)
}

func TestService_Search(t *testing.T) {

	repository := MockRepository{}

	repository.On("Search").Return([]Record{}, nil)

	service := ItemService{repository}

	items, _ := service.Search()

	for i := range items {
		assert.Equal(t, items[i].ID, int64(12), "item id true")
		assert.Equal(t, items[i].Title, "Apple", "item title true")
		assert.Equal(t, items[i].Content, "Apple Handphone", "item title true")
	}
}

// func TestService_SearchByCat(t *testing.T) {
// 	repository := ItemRepositoryMock{}
// 	repository.On("SearchByCat").Return([]Record{}, nil)
//
// 	service := ItemService{repository}
// 	items, _ := service.SearchByCat()
// 	for i := range items {
// 		assert.Equal(t, items[i].ID, int64(12), "item id true")
// 		assert.Equal(t, items[i].Title, "Apple", "item title true")
// 		assert.Equal(t, items[i].Content, "Apple Handphone", "item title true")
// 	}
// 	fmt.Println(items)
// }
