package service

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/model"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/framework/dal"
)

type PostService struct{}

// 创建岗位
func (s *PostService) CreatePost(param dto.SavePost) error {

	return dal.Gorm.Model(model.SysPost{}).Create(&model.SysPost{
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		Remark:   param.Remark,
		CreateBy: param.CreateBy,
	}).Error
}

// 删除岗位
func (s *PostService) DeletePost(postIds []int) error {
	return dal.Gorm.Model(model.SysPost{}).Where("post_id IN ?", postIds).Delete(&model.SysPost{}).Error
}

// 更新岗位
func (s *PostService) UpdatePost(param dto.SavePost) error {

	return dal.Gorm.Model(model.SysPost{}).Where("post_id = ?", param.PostId).Updates(&model.SysPost{
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		Remark:   param.Remark,
		UpdateBy: param.UpdateBy,
	}).Error
}

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

// 根据岗位id获取岗位详情
func (s *PostService) GetPostByPostId(postId int) dto.PostDetailResponse {

	var post dto.PostDetailResponse

	dal.Gorm.Model(model.SysPost{}).Where("post_id = ?", postId).Last(&post)

	return post
}

// 根据岗位名称获取岗位详情
func (s *PostService) GetPostByPostName(postName string) dto.PostDetailResponse {

	var post dto.PostDetailResponse

	dal.Gorm.Model(model.SysPost{}).Where("post_name = ?", postName).Last(&post)

	return post
}

// 根据岗位编码获取岗位详情
func (s *PostService) GetPostByPostCode(postCode string) dto.PostDetailResponse {

	var post dto.PostDetailResponse

	dal.Gorm.Model(model.SysPost{}).Where("post_code = ?", postCode).Last(&post)

	return post
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
