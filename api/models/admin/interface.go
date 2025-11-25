package apimodelsadmin

import "context"

type IModelAdmin interface {
	GetAdminDetail(ctx context.Context, request GetAdminDetailRequest) (response *GetAdminDetailResponse, err error)
}
