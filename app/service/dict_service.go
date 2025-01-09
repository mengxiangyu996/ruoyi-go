package service

import (
	"context"
	"encoding/json"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	rediskey "ruoyi-go/common/types/redis-key"
	"ruoyi-go/framework/dal"
)

type DictTypeService struct{}

// 创建字典类型
func (s *DictTypeService) CreateDictType(param dto.SaveDictType) error {

	return dal.Gorm.Model(model.SysDictType{}).Create(&model.SysDictType{
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		Remark:   param.Remark,
		CreateBy: param.CreateBy,
	}).Error
}

// 更新字典类型
func (s *DictTypeService) UpdateDictType(param dto.SaveDictType) error {

	return dal.Gorm.Model(model.SysDictType{}).Where("dict_id = ?", param.DictId).Updates(&model.SysDictType{
		DictName: param.DictName,
		DictType: param.DictType,
		Status:   param.Status,
		Remark:   param.Remark,
		UpdateBy: param.UpdateBy,
	}).Error
}

// 删除字典类型
func (s *DictTypeService) DeleteDictType(dictIds []int) error {
	return dal.Gorm.Model(model.SysDictType{}).Where("dict_id IN ?", dictIds).Delete(&model.SysDictType{}).Error
}

// 字典类型列表
func (s *DictTypeService) GetDictTypeList(param dto.DictTypeListRequest, isPaging bool) ([]dto.DictTypeListResponse, int) {

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

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&dictTypes)

	return dictTypes, int(count)
}

// 字典类型详情
func (s *DictTypeService) GetDictTypeByDictId(dictId int) dto.DictTypeDetailResponse {

	var dictType dto.DictTypeDetailResponse

	dal.Gorm.Model(model.SysDictType{}).Where("dict_id = ?", dictId).Last(&dictType)

	return dictType
}

// 根据字典类型查询详情
func (s *DictTypeService) GetDcitTypeByDictType(dictType string) dto.DictTypeDetailResponse {

	var dictTypeResult dto.DictTypeDetailResponse

	dal.Gorm.Model(model.SysDictType{}).Where("dict_type = ?", dictType).Last(&dictTypeResult)

	return dictTypeResult
}

type DictDataService struct{}

// 创建字典数据
func (s *DictDataService) CreateDictData(param dto.SaveDictData) error {

	return dal.Gorm.Model(model.SysDictData{}).Create(&model.SysDictData{
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		Remark:    param.Remark,
		CreateBy:  param.CreateBy,
	}).Error
}

// 更新字典数据
func (s *DictDataService) UpdateDictData(param dto.SaveDictData) error {

	return dal.Gorm.Model(model.SysDictData{}).Where("dict_code = ?", param.DictCode).Updates(&model.SysDictData{
		DictSort:  param.DictSort,
		DictLabel: param.DictLabel,
		DictValue: param.DictValue,
		DictType:  param.DictType,
		CssClass:  param.CssClass,
		ListClass: param.ListClass,
		IsDefault: param.IsDefault,
		Status:    param.Status,
		Remark:    param.Remark,
		UpdateBy:  param.UpdateBy,
	}).Error
}

// 删除字典数据
func (s *DictDataService) DeleteDictData(dictCodes []int) error {
	return dal.Gorm.Model(model.SysDictData{}).Where("dict_code IN ?", dictCodes).Delete(&model.SysDictData{}).Error
}

// 字典数据列表
func (s *DictDataService) GetDictDataList(param dto.DictDataListRequest, isPaging bool) ([]dto.DictDataListResponse, int) {

	var count int64
	dictDatas := make([]dto.DictDataListResponse, 0)

	query := dal.Gorm.Model(model.SysDictData{}).Order("dict_code")

	if param.DictLabel != "" {
		query = query.Where("dict_label LIKE ?", "%"+param.DictLabel+"%")
	}

	if param.DictType != "" {
		query = query.Where("dict_type LIKE ?", "%"+param.DictType+"%")
	}

	if param.Status != "" {
		query = query.Where("status = ?", param.Status)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&dictDatas)

	return dictDatas, int(count)
}

// 根据字典数据编码获取字典数据详情
func (s *DictDataService) GetDictDataByDictCode(dictCode int) dto.DictDataDetailResponse {

	var dictData dto.DictDataDetailResponse

	dal.Gorm.Model(model.SysDictData{}).Where("dict_code = ?", dictCode).Last(&dictData)

	return dictData
}

// 根据字典类型查询字典数据
func (s *DictDataService) GetDictDataByDictType(dictType string) []dto.DictDataListResponse {

	dictDatas := make([]dto.DictDataListResponse, 0)

	dal.Gorm.Model(model.SysDictData{}).Where("status = ? AND dict_type = ?", constant.NORMAL_STATUS, dictType).Find(&dictDatas)

	return dictDatas
}

// 根据字典类型查询字典数据
func (s *DictDataService) GetDictDataCacheByDictType(dictType string) []dto.DictDataListResponse {

	dictDatas := make([]dto.DictDataListResponse, 0)

	// 缓存不为空不从数据库读取，减少数据库压力
	if dictDatasCache, _ := dal.Redis.HGet(context.Background(), rediskey.SysDictKey, dictType).Result(); dictDatasCache != "" {
		if err := json.Unmarshal([]byte(dictDatasCache), &dictDatas); err == nil {
			return dictDatas
		}
	}

	// 从数据库读取配置并且记录到缓存
	dictDatas = s.GetDictDataByDictType(dictType)
	if len(dictDatas) > 0 {
		dictDadasBytes, _ := json.Marshal(&dictDatas)
		dal.Redis.HSet(context.Background(), rediskey.SysDictKey, dictType, string(dictDadasBytes))
	}

	return dictDatas
}
