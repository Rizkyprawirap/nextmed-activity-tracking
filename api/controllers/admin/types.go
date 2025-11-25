package apicontrollersadmin

import apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"

type (
	GetAdminDetailRequest struct {
		Email    string
		Password string
	}

	GetAdminDetailResponse struct {
		Token string       `json:"token"`
		Data  apidto.Admin `json:"data"`
	}
)
