package apicontrollersusage

import "context"

type (
	IControllerUsage interface {
		GetDailyUsage(ctx context.Context, request GetDailyUsageRequest) (response *GetDailyUsageResponse, err error)
		GetTopClientUsage(ctx context.Context, request GetTopClientUsageRequest) (response *GetTopClientUsageResponse, err error)
	}
)
