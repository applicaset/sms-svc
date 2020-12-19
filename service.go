package smssvc

import "context"

type Service interface {
	SendMessage(ctx context.Context, phoneNumber, message string) (err error)
}
