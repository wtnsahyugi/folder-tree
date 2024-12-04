package usecase

import (
	"context"
	"folder-tree/internal/entity"
	repository "folder-tree/pkg/repository/postgres"
)

func (f *folderImpl) GetFolderChild(ctx context.Context, id int64, depth int32) (entity.TreeFolder, error) {
	folders, err := f.repo.ListFolderTreeById(ctx, repository.ListFolderTreeByIdParams{
		ID:    id,
		Depth: depth,
	})
	if err != nil {
		return entity.TreeFolder{}, err
	}

	if len(folders) == 0 {
		return entity.TreeFolder{}, nil
	}

	root := entity.TreeFolder{
		ID:       int64(folders[0].ID),
		Name:     folders[0].Name,
		IsFolder: folders[0].IsFolder.Bool,
		ParentID: int64(folders[0].ParentID.Int32),
	}
	mapChild := make(map[int64][]entity.TreeFolder)
	for _, f := range folders[1:] {
		mapChild[int64(f.ParentID.Int32)] = append(mapChild[int64(f.ParentID.Int32)], entity.TreeFolder{
			ID:       int64(f.ID),
			Name:     f.Name,
			IsFolder: f.IsFolder.Bool,
			ParentID: int64(f.ParentID.Int32),
		})
	}

	child := mapChild[int64(root.ID)]
	root = construct(root, child, mapChild)

	if root.ID != 0 {
		return root, nil
	}

	return entity.TreeFolder{}, nil
}

func construct(parent entity.TreeFolder, children []entity.TreeFolder, mapChild map[int64][]entity.TreeFolder) entity.TreeFolder {
	if len(children) == 0 {
		return parent
	}

	for _, c := range children {
		parent.Child = append(parent.Child, construct(c, mapChild[c.ID], mapChild))
	}

	return parent
}

func (f *folderImpl) GetAllChild(ctx context.Context) (entity.TreeFolder, error) {
	return entity.TreeFolder{}, nil
}
