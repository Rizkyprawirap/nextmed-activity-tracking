package apicontrollersadmin

import (
	"context"

	apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"
	apimodelsadmin "github.com/Rizkyprawirap/nextmed-activity-tracking/api/models/admin"
	pkgJWT "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type controller struct {
	modelAdmin apimodelsadmin.IModelAdmin
	jwt        pkgJWT.IPkgJWT
}

func New(
	modelAdmin apimodelsadmin.IModelAdmin,
	jwt pkgJWT.IPkgJWT,
) IControllerAdmin {
	return &controller{
		modelAdmin: modelAdmin,
		jwt:        jwt,
	}
}

func (c *controller) Login(ctx context.Context, request GetAdminDetailRequest) (*GetAdminDetailResponse, error) {

	response, err := c.modelAdmin.GetAdminDetail(ctx, apimodelsadmin.GetAdminDetailRequest{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(response.Data.Password), []byte(request.Password)); err != nil {
	}

	token, err := c.jwt.Generate(ctx, response.Data)
	if err != nil {
		return nil, err
	}

	return &GetAdminDetailResponse{
		Token: token,
		Data: apidto.Admin{
			ID:        response.Data.ID,
			Email:     response.Data.Email,
			FullName:  response.Data.FullName,
			CreatedAt: response.Data.CreatedAt,
			UpdatedAt: response.Data.UpdatedAt,
		},
	}, nil
}
