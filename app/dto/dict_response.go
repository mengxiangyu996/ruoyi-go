package dto

import "ruoyi-go/framework/datetime"

// 字典类型列表
type DictTypeListResponse struct {
	DictId     int               `json:"dictId"`
	DictName   string            `json:"dictName"`
	DictType   string            `json:"dictType"`
	Status     string            `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
	Remark     string            `json:"remark"`
}

// 字典类型详情
type DictTypeDetailResponse struct {
	DictId   int    `json:"dictId"`
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

// 字典数据列表
type DictDataListResponse struct {
	DictCode   int               `json:"dictCode"`
	DictSort   int               `json:"dictSort"`
	DictLabel  string            `json:"dictLabel"`
	DictValue  string            `json:"dictValue"`
	DictType   string            `json:"dictType"`
	CssClass   string            `json:"cssClass"`
	ListClass  string            `json:"listClass"`
	IsDefault  string            `json:"isDefault"`
	Status     string            `json:"status"`
	CreateTime datetime.Datetime `json:"createTime"`
	Default    bool              `json:"default" gorm:"-"`
}

// 字典数据详情
type DictDataDetailResponse struct {
	DictCode  int    `json:"dictCode"`
	DictSort  int    `json:"dictSort"`
	DictLabel string `json:"dictLabel"`
	DictValue string `json:"dictValue"`
	DictType  string `json:"dictType"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault string `json:"isDefault"`
	Status    string `json:"status"`
	Default   bool   `json:"default" gorm:"-"`
}
