package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type LogininforService struct{}

// 删除登录日志
func (s *LogininforService) DeleteLogininfor(infoIds []int) error {

	if len(infoIds) > 0 {
		return dal.Gorm.Model(model.SysLogininfor{}).Where("info_id IN ?", infoIds).Delete(&model.SysLogininfor{}).Error
	}

	// 为解决 WHERE conditions required 错误，添加 Where("info_id > ?", 0) 这个条件
	return dal.Gorm.Model(model.SysLogininfor{}).Where("info_id > ?", 0).Delete(&model.SysLogininfor{}).Error
}

// 获取登录日志列表
func (s *LogininforService) GetLogininforList(param dto.LogininforListRequest, isPaging bool) ([]dto.LogininforListResponse, int) {

	var count int64
	logininfos := make([]dto.LogininforListResponse, 0)

	query := dal.Gorm.Model(model.SysLogininfor{}).Order(param.OrderByColumn + " " + param.OrderRule)

	if param.Ipaddr != "" {
		query = query.Where("ipaddr LIKE ?", "%"+param.Ipaddr+"%")
	}

	if param.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+param.UserName+"%")
	}

	if param.Status != "" {
		query = query.Where("status = ?", param.Status)
	}

	if param.BeginTime != "" && param.EndTime != "" {
		query = query.Where("login_time BETWEEN ? AND ?", param.BeginTime, param.EndTime)
	}

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&logininfos)

	return logininfos, int(count)
}

// 记录登录信息
func (s *LogininforService) CreateSysLogininfor(param dto.SaveLogininforRequest) error {

	go func() error {
		return dal.Gorm.Model(model.SysLogininfor{}).Create(&model.SysLogininfor{
			UserName:      param.UserName,
			Ipaddr:        param.Ipaddr,
			LoginLocation: param.LoginLocation,
			Browser:       param.Browser,
			Os:            param.Os,
			Status:        param.Status,
			Msg:           param.Msg,
			LoginTime:     param.LoginTime,
		}).Error
	}()

	return nil
}
