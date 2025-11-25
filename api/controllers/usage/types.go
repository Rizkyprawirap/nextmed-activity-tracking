package apicontrollersusage

import apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"

type (
	GetDailyUsageRequest struct {
		APIKey string `form:"api_key"`
	}
	GetDailyUsageResponse     []apidto.Usage
	GetTopClientUsageRequest  struct{}
	GetTopClientUsageResponse []apidto.TopClientUsage
)
