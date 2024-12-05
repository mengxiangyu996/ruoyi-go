package dto

type PageRequest struct {
	PageNum  int `query:"pageNum" form:"pageNum"`
	PageSize int `query:"pageSize" form:"pageSize"`
}
