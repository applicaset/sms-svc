package sms_svc

import (
	"context"
	"github.com/pkg/errors"
)

type service struct {
	sender MessageSender
}

func (svc *service) SendMessage(ctx context.Context, phoneNumber, message string) error {
	err := svc.sender.Send(ctx, phoneNumber, message)
	if err != nil {
		return errors.Wrap(err, "error on send message")
	}

	return nil
}

type MessageSender interface {
	Send(ctx context.Context, phoneNumber, message string) (err error)
}

func New(sender MessageSender) Service {
	svc := service{
		sender: sender,
	}

	return &svc
}
