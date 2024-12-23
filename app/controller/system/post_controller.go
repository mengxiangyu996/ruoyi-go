package systemcontroller

import (
	"ruoyi-go/app/controller/validator"
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/app/token"
	"ruoyi-go/common/types/constant"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

// 岗位列表
func (*PostController) List(ctx *gin.Context) {

	var param dto.PostListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	posts, total := (&service.PostService{}).GetPostList(param, true)

	response.NewSuccess().SetPageData(posts, total).Json(ctx)
}

// 岗位详情
func (*PostController) Detail(ctx *gin.Context) {

	postId, _ := strconv.Atoi(ctx.Param("postId"))

	post := (&service.PostService{}).GetPostByPostId(postId)

	response.NewSuccess().SetData("data", post).Json(ctx)
}

// 新增岗位
func (*PostController) Create(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_INSERT)

	var param dto.CreatePostRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreatePostValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if post := (&service.PostService{}).GetPostByPostName(param.PostName); post.PostId > 0 {
		response.NewError().SetMsg("新增岗位" + param.PostName + "失败，岗位名称已存在").Json(ctx)
		return
	}

	if post := (&service.PostService{}).GetPostByPostCode(param.PostCode); post.PostId > 0 {
		response.NewError().SetMsg("新增岗位" + param.PostName + "失败，岗位编码已存在").Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.PostService{}).CreatePost(dto.SavePost{
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		CreateBy: loginUser.UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新岗位
func (*PostController) Update(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_UPDATE)

	var param dto.UpdatePostRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdatePostValidator(param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if post := (&service.PostService{}).GetPostByPostName(param.PostName); post.PostId > 0 && post.PostId != param.PostId {
		response.NewError().SetMsg("修改岗位" + param.PostName + "失败，岗位名称已存在").Json(ctx)
		return
	}

	if post := (&service.PostService{}).GetPostByPostCode(param.PostCode); post.PostId > 0 && post.PostId != param.PostId {
		response.NewError().SetMsg("修改岗位" + param.PostName + "失败，岗位编码已存在").Json(ctx)
		return
	}

	loginUser, _ := token.GetLoginUser(ctx)

	if err := (&service.PostService{}).UpdatePost(dto.SavePost{
		PostId:   param.PostId,
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		UpdateBy: loginUser.UserName,
		Remark:   param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除岗位
func (*PostController) Remove(ctx *gin.Context) {

	// 设置业务类型，操作日志获取
	ctx.Set(constant.REQUEST_BUSINESS_TYPE, constant.REQUEST_BUSINESS_TYPE_DELETE)

	postIds, err := utils.StringToIntSlice(ctx.Param("postIds"), ",")
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	if err = (&service.PostService{}).DeletePost(postIds); err != nil {
		response.NewError().SetMsg(err.Error())
		return
	}

	response.NewSuccess().Json(ctx)
}
