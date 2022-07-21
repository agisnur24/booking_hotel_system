package controller

import "context"

type LoginController interface {
	Login(ctx context.Context) error
}
