package apicontrollersadmin

import "context"

type (
	IControllerAdmin interface {
		Login(ctx context.Context, request GetAdminDetailRequest) (response *GetAdminDetailResponse, err error)
	}
)
