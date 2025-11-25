package apimodelsusage

import (
	"context"
)

type IModelUsage interface {
	GetDailyUsage(ctx context.Context, request GetDailyUsageRequest) (response *GetDailyUsageResponse, err error)
	GetTopClientUsage(ctx context.Context) (response *GetTopClientUsageResponse, err error)
}
