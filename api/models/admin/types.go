package apimodelsadmin

import apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"

type (
	GetAdminDetailRequest struct {
		Email    string
		Password string
	}

	GetAdminDetailResponse struct {
		Data apidto.Admin
	}
)
