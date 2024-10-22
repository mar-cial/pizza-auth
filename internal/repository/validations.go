package repository

import "context"

type AuthValidations interface {
	PhonenumberExists(ctx context.Context, phonenumber string) error
}
