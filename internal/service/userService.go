package service

import (
	"GoCart/config"
	"GoCart/internal/domain"
	"GoCart/internal/dto"
	"GoCart/internal/helper"
	"GoCart/internal/repository"
	"GoCart/pkg/notification"
	"errors"
	"fmt"
	"log"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform some db operations
	// write buisness logic
	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	hPassword, err := s.Auth.CreateHashPassword(input.Password)
	if err != nil {
		return "", err
	}

	log.Println("user phone number getting is", input.Phone)
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) Login(email, password string) (string, error) {
	user, err := s.Repo.FindUser(email)
	if err != nil {
		return "", err
	}
	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}
	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) isVerifiedUser(id uint) bool {

	currentUser, err := s.Repo.FindUserById(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) error {

	if s.isVerifiedUser(e.ID) {
		return errors.New("user already verified")
	}

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return err
	}
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}
	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return errors.New("unable to update verification code")
	}
	user, _ = s.Repo.FindUserById(e.ID)

	//send sms
	notificationClient := notification.NewNotificationClient(s.Config)
	msg := fmt.Sprintf("Your verification code is %v", code)
	var toList []string
	toList = append(toList, user.Email)
	err = notificationClient.SendSms(toList, msg)
	if err != nil {
		return errors.New("error sending sms")
	}

	return nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	if s.isVerifiedUser(id) {
		return errors.New("user already verified")
	}
	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return err
	}
	if code != user.Code || !time.Now().Before(user.Expiry) {
		return errors.New("please provide valid code")
	}

	updateUser := domain.User{
		Verified: true,
	}
	_, err = s.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("unable to verify user")
	}

	return nil
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

func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {
	user, _ := s.Repo.FindUserById(id)

	if user.UserType == domain.SELLER {
		return "", errors.New("user is already seller")
	}

	_, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	})
	if err != nil {
		return "", errors.New("user unable to become seller")
	}

	token, err := s.Auth.GenerateToken(id, user.Email, domain.SELLER)
	if err != nil {
		return "", err
	}

	account := domain.BankAccount{
		BankAccount: input.BankAccountNumber,
		IFSCCode:    input.IFSCCode,
		PaymentType: input.PaymentType,
		UserId:      id,
	}
	err = s.Repo.CreateBankAccount((account))
	if err != nil {
		return "", err
	}

	return token, nil
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
