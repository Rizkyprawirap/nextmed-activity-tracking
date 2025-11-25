package apimodelsadmin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"
	pkgpostgre "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/postgre"
)

type model struct {
	pkgPostgre pkgpostgre.IPkgPostgreDB
}

func New(
	pkgPostgre pkgpostgre.IPkgPostgreDB,
) IModelAdmin {
	return &model{
		pkgPostgre: pkgPostgre,
	}
}

func (m *model) GetAdminDetail(ctx context.Context, request GetAdminDetailRequest) (response *GetAdminDetailResponse, err error) {
	qctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, password, full_name, created_at, updated_at
		FROM admins
		WHERE deleted_at IS NULL
		AND email = $1
		LIMIT 1
	`

	var adm apidto.Admin

	err = m.pkgPostgre.GetConnection().
		QueryRowContext(qctx, query, request.Email).
		Scan(
			&adm.ID,
			&adm.Email,
			&adm.Password,
			&adm.FullName,
			&adm.CreatedAt,
			&adm.UpdatedAt,
		)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("super admin not found")
		}
		return nil, fmt.Errorf("failed to query super admin: %w", err)
	}

	return &GetAdminDetailResponse{Data: adm}, nil
}
