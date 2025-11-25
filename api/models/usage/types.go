package apimodelsusage

import apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"

type (
	GetDailyUsageRequest struct {
		APIKey string
	}
	GetDailyUsageResponse     []apidto.Usage
	GetTopClientUsageResponse []apidto.TopClientUsage
)
