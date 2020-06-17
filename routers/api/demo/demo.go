package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/knight0zh/demo_server/base"
	"github.com/knight0zh/demo_server/service"
)

func HelloWorld(c *gin.Context) {
	ctx := base.NewContext(c)

	uId := ctx.Value("user_id")
	var params struct {
		BAutoId string `form:"b_auto_id" binding:"required"`
		StoreId int    `form:"store_id" binding:"lt=2"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.Alert(1000, err.Error())
		return
	}

	service.DemoService()
	ctx.Success(map[string]interface{}{"data": "hello world", "user_id": uId})
}
