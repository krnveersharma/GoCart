package service

import (
	"GoCart/internal/domain"
	"GoCart/internal/dto"
	"log"
)

type UserService struct {
}

func (s UserService) findUserByEmail(email string) (domain.User, error) {
	// perform some db operations
	// write buisness logic
	return domain.User{}, nil
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println(input)

	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) (int, error) {

	return 0, nil
}

func (s UserService) CreateProfile(id uint, input interface{}) error {

	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {

	return nil
}

func (s UserService) BecomeSeller(id uint, input interface{}) (string, error) {

	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateCart(input interface{}, u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (interface{}, error) {

	return nil, nil
}
