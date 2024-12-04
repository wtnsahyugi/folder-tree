package entity

type TreeFolder struct {
	ID       int64        `json:"id"`
	Name     string       `json:"name"`
	IsFolder bool         `json:"is_folder"`
	ParentID int64        `json:"parent_id"`
	Child    []TreeFolder `json:"child"`
}
