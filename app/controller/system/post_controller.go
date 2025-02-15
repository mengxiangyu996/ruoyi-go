package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/security"
	"ruoyi-go/app/service"
	"ruoyi-go/app/validator"
	"ruoyi-go/common/utils"
	"ruoyi-go/framework/response"
	"strconv"
	"time"

	"gitee.com/hanshuangjianke/go-excel/excel"
	"github.com/gin-gonic/gin"
)

type PostController struct{}

// 岗位列表
func (*PostController) List(ctx *gin.Context) {

	var param dto.PostListRequest

	if err := ctx.ShouldBind(&param); err != nil {
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

	var param dto.CreatePostRequest

	if err := ctx.ShouldBind(&param); err != nil {
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

	if err := (&service.PostService{}).CreatePost(dto.SavePost{
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		CreateBy: security.GetAuthUserName(ctx),
		Remark:   param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新岗位
func (*PostController) Update(ctx *gin.Context) {

	var param dto.UpdatePostRequest

	if err := ctx.ShouldBind(&param); err != nil {
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

	if err := (&service.PostService{}).UpdatePost(dto.SavePost{
		PostId:   param.PostId,
		PostCode: param.PostCode,
		PostName: param.PostName,
		PostSort: param.PostSort,
		Status:   param.Status,
		UpdateBy: security.GetAuthUserName(ctx),
		Remark:   param.Remark,
	}); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除岗位
func (*PostController) Remove(ctx *gin.Context) {

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

// 数据导出
func (*PostController) Export(ctx *gin.Context) {

	var param dto.PostListRequest

	if err := ctx.ShouldBind(&param); err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	list := make([]dto.PostExportResponse, 0)

	posts, _ := (&service.PostService{}).GetPostList(param, false)
	for _, post := range posts {
		list = append(list, dto.PostExportResponse{
			PostId:   post.PostId,
			PostCode: post.PostCode,
			PostName: post.PostName,
			PostSort: post.PostSort,
			Status:   post.Status,
		})
	}

	file, err := excel.NormalDynamicExport("Sheet1", "", "", false, false, list, nil)
	if err != nil {
		response.NewError().SetMsg(err.Error()).Json(ctx)
		return
	}

	excel.DownLoadExcel("post_"+time.Now().Format("20060102150405"), ctx.Writer, file)
}
