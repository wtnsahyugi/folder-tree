package handler

import (
	"folder-tree/internal/usecase"
	repository "folder-tree/pkg/repository/postgres"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type folderHandler struct {
	uc usecase.FolderUseCase
}

type FolderHandlerIface interface {
	GetFolderChild(ctx *gin.Context)
}

func NewFolderHandler(db *pgxpool.Pool) FolderHandlerIface {
	return &folderHandler{usecase.NewFolderUseCase(repository.New(db))}
}

func (h *folderHandler) GetFolderChild(ctx *gin.Context) {
	parentID := ctx.Query("id")
	id, err := strconv.Atoi(parentID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	depth := ctx.Query("depth")
	depthInt, err := strconv.Atoi(depth)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid depth"})
		return
	}

	result, err := h.uc.GetFolderChild(ctx, int64(id), int32(depthInt))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to get folder child"})
		return
	}

	ctx.JSON(200, gin.H{
		"data":    result,
		"message": "success",
	})
}
