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

// 岗位详情
type PostDetailResponse struct {
	PostId   int    `json:"postId"`
	PostCode string `json:"postCode"`
	PostName string `json:"postName"`
	PostSort int    `json:"postSort"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

// 岗位导出
type PostExportResponse struct {
	PostId   int    `excel:"name:岗位序号;"`
	PostCode string `excel:"name:岗位编码;"`
	PostName string `excel:"name:岗位名称;"`
	PostSort int    `excel:"name:岗位排序;"`
	Status   string `excel:"name:状态;replace:0_正常,1_停用;"`
}
