package dto

// 下拉选项树
type SeleteTree struct {
	Id       int          `json:"id"`
	Label    string       `json:"label"`
	Children []SeleteTree `json:"children" gorm:"-"`
	ParentId int          `json:"-"`
}
