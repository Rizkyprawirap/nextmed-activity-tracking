package apimodelsclient

import "context"

type IModelClient interface {
	InsertClient(ctx context.Context, request InsertClientRequest) (*InsertClientResponse, error)
	InsertLog(ctx context.Context, request InsertLogRequest) error
	GetClientByAPIKey(ctx context.Context, request GetClientByAPIKeyRequest) (response *GetClientByAPIKeyResponse, err error)
}
