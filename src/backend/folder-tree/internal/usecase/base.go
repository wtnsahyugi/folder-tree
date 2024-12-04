package usecase

import (
	"context"
	"folder-tree/internal/entity"
	repository "folder-tree/pkg/repository/postgres"
)

type folderImpl struct {
	repo repository.Querier
}

type FolderUseCase interface {
	GetFolderChild(ctx context.Context, id int64, depth int32) (entity.TreeFolder, error)
	GetAllChild(ctx context.Context) (entity.TreeFolder, error)
}

func NewFolderUseCase(repo repository.Querier) FolderUseCase {
	return &folderImpl{
		repo: repo,
	}
}
