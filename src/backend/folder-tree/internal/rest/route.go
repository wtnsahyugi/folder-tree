package rest

import (
	"folder-tree/internal/handler"
	"folder-tree/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func registerRoute(router *gin.Engine, db *pgxpool.Pool) {
	if router == nil {
		panic("router is not instantiated")
	}

	h := handler.NewFolderHandler(db)

	routes := router.Group("/", middleware.CORSMiddleware())
	{
		routes.GET("folder/tree", h.GetFolderChild)
	}
}
