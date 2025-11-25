package pkgjwt

import apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"

type (
	ValidateRequest struct {
		Token string
	}
	ValidateResponse struct {
		Claims apidto.JWTClaims
	}
	AdminData struct {
		ID string `json:"id"`
	}
)
