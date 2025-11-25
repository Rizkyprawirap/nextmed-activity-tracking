package apicontrollersclient

import "context"

type (
	IControllerClient interface {
		CreateClient(ctx context.Context, request CreateClientRequest) (response *CreateClientResponse, err error)
		InsertLog(ctx context.Context, request InsertLogRequest) error
	}
)
