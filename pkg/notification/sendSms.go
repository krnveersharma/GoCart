package notification

import (
	"GoCart/config"
	"errors"
	"fmt"
	"net/smtp"
)

type NotificationClient interface {
	SendSms(phone []string, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

func (c notificationClient) SendSms(to []string, message string) error {
	host := "smtp.gmail.com"
	port := "587"
	body := []byte(message)
	auth := smtp.PlainAuth("", c.config.Email, c.config.Password, host)
	err := smtp.SendMail(host+":"+port, auth, c.config.Email, to, body)

	if err != nil {
		fmt.Errorf("Error in Sending Email %v", err.Error())
		return errors.New("unable to send email")
	}
	return nil
}

func NewNotificationClient(config config.AppConfig) NotificationClient {

	return &notificationClient{
		config: config,
	}
}
