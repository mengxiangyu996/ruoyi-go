package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/framework/dal"
)

type PostService struct{}

// 岗位列表
func (s *PostService) GetPostList(param dto.PostListRequest) ([]dto.PostListResponse, int) {

	var count int64
	posts := make([]dto.PostListResponse, 0)

	query := dal.Gorm.Model(model.SysPost{}).Order("post_sort, post_id")

	if param.PostCode != "" {
		query.Where("post_code LIKE ?", "%"+param.PostCode+"%")
	}

	if param.PostName != "" {
		query.Where("post_name LIKE ?", "%"+param.PostName+"%")
	}

	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}

	query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&posts)

	return posts, int(count)
}