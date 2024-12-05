package dto

// 岗位列表
type PostListRequest struct {
	PageRequest
	PostCode string `query:"postCode" form:"postCode"`
	PostName string `query:"postName" form:"postName"`
	Status   string `query:"status" form:"status"`
}
