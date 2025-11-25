package apicontrollersusage

import (
	"context"
	"fmt"

	apimodelsusage "github.com/Rizkyprawirap/nextmed-activity-tracking/api/models/usage"
	pkgJWT "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/jwt"
)

type controller struct {
	modelsUsage apimodelsusage.IModelUsage
	jwt         pkgJWT.IPkgJWT
}

func New(
	modelsUsage apimodelsusage.IModelUsage,
	jwt pkgJWT.IPkgJWT,
) IControllerUsage {
	return &controller{
		modelsUsage: modelsUsage,
		jwt:         jwt,
	}
}

func (c *controller) GetDailyUsage(
	ctx context.Context,
	request GetDailyUsageRequest,
) (response *GetDailyUsageResponse, err error) {

	usageData, err := c.modelsUsage.GetDailyUsage(ctx, apimodelsusage.GetDailyUsageRequest{
		APIKey: request.APIKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed getting daily usage: %w", err)
	}

	response = (*GetDailyUsageResponse)(usageData)

	return response, nil
}
func (c *controller) GetTopClientUsage(
	ctx context.Context,
	request GetTopClientUsageRequest,
) (response *GetTopClientUsageResponse, err error) {

	topUsageClientData, err := c.modelsUsage.GetTopClientUsage(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting daily usage: %w", err)
	}

	response = (*GetTopClientUsageResponse)(topUsageClientData)

	return response, nil
}
