package routesusage

type (
	GetDailyUsageRequest struct {
		APIKey string `form:"api_key" binding:"required"`
	}
)
