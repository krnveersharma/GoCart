package notification

import "GoCart/config"

type NotificationClient interface {
	SendSms(phone string, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

func (c notificationClient) SendSms(phone string, message string) error {

	return nil
}

func NewNotificationClient(config config.AppConfig) NotificationClient {

	return &notificationClient{
		config: config,
	}
}
