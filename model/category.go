package model

type Category struct {
	ID                  uint64 `json:"id"`
	Name                string `json:"name"`
	ReviewableByGroupID uint64 `json:"reviewable_by_group_id"`
}
