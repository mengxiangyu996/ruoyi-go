package dto

import "ruoyi-go/framework/datetime"

// 岗位列表
type PostListResponse struct {
	PostId     int               `json:"postId"`
	PostCode   string            `json:"postCode"`
	PostName   string            `json:"postName"`
	PostSort   int               `json:"postSort"`
	Status     string            `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
}
