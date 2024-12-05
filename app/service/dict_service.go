package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type DictService struct{}

// 字典类型列表
func (s *DictService) GetDictTypeList(param dto.DictTypeRequest) ([]dto.DictTypeListResponse, int) {

	var count int64
	dictTypes := make([]dto.DictTypeListResponse, 0)

	query := dal.Gorm.Model(model.SysDictType{}).Order("dict_id")

	if param.DictName != "" {
		query = query.Where("dict_name LIKE ?", "%"+param.DictName+"%")
	}

	if param.DictType != "" {
		query = query.Where("dict_type LIKE ?", "%"+param.DictType+"%")
	}

	if param.Status != "" {
		query = query.Where("status = ?", param.Status)
	}

	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("create_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}

	query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&dictTypes)

	return dictTypes, int(count)
}

// 根据字典类型查询字典数据
func (s *DictService) GetDictDataByType(dictType string) []dto.DictDataListResponse {

	dictDatas := make([]dto.DictDataListResponse, 0)

	dal.Gorm.Model(model.SysDictData{}).Where("status = ? AND dict_type = ?", constant.NORMAL_STATUS, dictType).Find(&dictDatas)

	return dictDatas
}
