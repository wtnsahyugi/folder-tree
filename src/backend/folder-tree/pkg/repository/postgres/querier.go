// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package repository

import (
	"context"
)

type Querier interface {
	ListAllFolders(ctx context.Context) ([]ListAllFoldersRow, error)
	ListFolderTreeById(ctx context.Context, arg ListFolderTreeByIdParams) ([]ListFolderTreeByIdRow, error)
}

var _ Querier = (*Queries)(nil)
