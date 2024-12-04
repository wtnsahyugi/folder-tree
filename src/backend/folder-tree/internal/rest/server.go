package rest

import (
	"folder-tree/config"
	"folder-tree/pkg/pgx"

	"github.com/gin-gonic/gin"
)

func StartServer(cfg config.Config) error {
	pgxPool, err := pgx.NewPgx(cfg)
	if err != nil {
		return err
	}

	router := gin.New()
	registerRoute(router, pgxPool)
	return router.Run(cfg.ServerAddr)
}
