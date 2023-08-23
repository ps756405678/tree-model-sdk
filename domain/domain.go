package domain

import "time"

type TreeNode struct {
	ID         string    `json:"_id"`
	Title      string    `json:"title"`
	ParentId   string    `json:"parent_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
