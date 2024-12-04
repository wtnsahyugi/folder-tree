-- name: ListFolderTreeById :many
WITH RECURSIVE folder_hierarchy AS (
    SELECT id, name, parent_id, is_folder, 1 AS depth
    FROM tree_folders
    WHERE (CASE WHEN @id::bigint = 0 THEN parent_id IS NULL ELSE id = @id::bigint END)
    UNION ALL
    SELECT f.id, f.name, f.parent_id,f.is_folder, fh.depth + 1
    FROM tree_folders f
    INNER JOIN folder_hierarchy fh ON f.parent_id = fh.id
    WHERE fh.depth < @depth::int
)
SELECT id, name, parent_id, depth, is_folder FROM folder_hierarchy order by depth;

-- name: ListAllFolders :many
WITH RECURSIVE folder_hierarchy AS (
    SELECT id, name, parent_id, is_folder, 1 AS depth
    FROM tree_folders WHERE parent_id IS NULL
    UNION ALL
    SELECT f.id, f.name, f.parent_id,f.is_folder, fh.depth + 1
    FROM tree_folders f
    INNER JOIN folder_hierarchy fh ON f.parent_id = fh.id
)
SELECT id, name, parent_id, depth, is_folder FROM folder_hierarchy order by depth;
