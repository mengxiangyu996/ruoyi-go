package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type PostService struct{}

// 岗位列表
func (s *PostService) GetPostList(param dto.PostListRequest, isPaging bool) ([]dto.PostListResponse, int) {

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

	if isPaging {
		query.Count(&count).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize)
	}

	query.Find(&posts)

	return posts, int(count)
}

// 根据用户id查询岗位id集合
func (s *PostService) GetPostIdsByUserId(userId int) []int {

	var postIds []int

	dal.Gorm.Model(model.SysPost{}).
		Joins("JOIN sys_user_post ON sys_user_post.post_id = sys_post.post_id").
		Where("sys_user_post.user_id = ? AND sys_post.status = ?", userId, constant.NORMAL_STATUS).
		Pluck("sys_post.post_id", &postIds)

	return postIds

}

// 根据用户id查询角色名
func (s *PostService) GetPostNamesByUserId(userId int) []string {

	var postNames []string

	dal.Gorm.Model(model.SysPost{}).
		Joins("JOIN sys_user_post ON sys_user_post.post_id = sys_post.post_id").
		Where("sys_user_post.user_id = ? AND sys_post.status = ?", userId, constant.NORMAL_STATUS).
		Pluck("sys_post.post_name", &postNames)

	return postNames
}
