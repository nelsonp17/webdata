package sqlc

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nelsonp17/webdata/app/database/sqlc/schemas"
)

type Repo struct {
	Queries *schemas.Queries
	Pgx     *pgxpool.Pool
}
