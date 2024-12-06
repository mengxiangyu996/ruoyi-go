package systemcontroller

import (
	"ruoyi-go/app/dto"
	"ruoyi-go/app/service"
	"ruoyi-go/framework/response"

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
