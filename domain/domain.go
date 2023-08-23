package domain

import "time"

type TreeNode struct {
	ID         string    `json:"_id"`
	Title      string    `json:"title"`
	ParentId   string    `json:"parent_id"`
	Data       any       `json:"data"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type Result struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
	Data       any    `json:"data"`
}
