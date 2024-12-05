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
		response.NewError().SetMsg(err.Error()).ToJson(ctx)
		return
	}

	posts, total := (&service.PostService{}).GetPostList(param)

	response.NewSuccess().SetPageData(posts, total).ToJson(ctx)
}
