package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nelsonp17/webdata/app/database/sqlc"
	"github.com/nelsonp17/webdata/app/database/sqlc/schemas"
	"github.com/nelsonp17/webdata/app/router/api"
)

type Router struct {
	// Fwk web framework.
	Fwk *fiber.App
}

type Handler struct {
	Database *pgxpool.Pool
}

func (h *Handler) Api(r Router) {
	route := r.Fwk.Group("/api/v1")

	apiRoute := api.Handler{
		Repo: sqlc.Repo{Queries: schemas.New(h.Database), Pgx: h.Database},
	}

	route.Get("/change_dollar", apiRoute.GetPriceDollar)
}
