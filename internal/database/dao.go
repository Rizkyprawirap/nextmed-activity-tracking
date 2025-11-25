package internaldatabase

import (
	"time"

	"github.com/google/uuid"
)

type (
	Client struct {
		ID        uuid.UUID  `db:"client_id"`
		Name      string     `db:"name"`
		Email     string     `db:"email"`
		ApiKey    string     `db:"api_key"`
		CreatedAt time.Time  `db:"created_at"`
		UpdatedAt time.Time  `db:"updated_at"`
		DeletedAt *time.Time `db:"deleted_at"`
	}
)
