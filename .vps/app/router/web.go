package router

import (
	"github.com/nelsonp17/webdata/app/database/sqlc"
	"github.com/nelsonp17/webdata/app/database/sqlc/schemas"
	"github.com/nelsonp17/webdata/app/router/web"
)

func (h *Handler) Web(r Router) {
	route := r.Fwk.Group("/")

	webRoute := web.Handler{
		Repo: sqlc.Repo{Queries: schemas.New(h.Database), Pgx: h.Database},
	}

	route.Get("/", webRoute.HomeView)
	route.Get("/assets/download", webRoute.DownloadApp)
}
