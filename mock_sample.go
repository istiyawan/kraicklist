package main

import "fmt"

// Model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Interface
type UserRepositoryInterface interface {
	GetAllUsers() ([]User, error)
}

// Service
type UserService struct {
	UserRepositoryInterface
}
func (s UserService) GetUser() ([]User, error){
	users, _ := s.UserRepositoryInterface.GetAllUsers()
	for i := range users {
		users[i].Password = "*****"
	}
	return users, nil
}

// Repository
type UserRepository struct {}
func (r UserRepository) GetAllUsers() ([]User, error) {
	users := []User{
		{"real", "real"},
		{"real2", "real2"},
	}
	return users, nil
}

// Main
func main__(){
	repository := UserRepository{}
	service := UserService{repository}
	users, _ := service.GetUser()
	fmt.Println(users)
}