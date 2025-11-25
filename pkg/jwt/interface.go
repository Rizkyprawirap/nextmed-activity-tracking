package pkgjwt

import "context"

type IPkgJWT interface {
	Generate(ctx context.Context, payload any) (string, error)
	Revoke(ctx context.Context, token string) error
	Validate(ctx context.Context, request ValidateRequest) (response *ValidateResponse, err error)
}
