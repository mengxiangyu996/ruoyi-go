package dto

type MenuListRequest struct {
	MenuName string `query:"menuName" form:"menuName"`
	Status   string `query:"status" form:"status"`
}
