package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/knight0zh/demo_pkg/pkg"
	"github.com/knight0zh/demo_server/base"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 10000,
				"msg":  "token error",
				"data": nil,
			})
			c.Abort()
			return
		}

		var response map[string]interface{}
		if err := pkg.HttpGet("user.com/user/userinfor").Param("token", token).Param("av", "1").
			ToJSON(&response); err != nil {
			base.ErrLog(err)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  "system error",
				"data": nil,
			})
			c.Abort()
			return
		}
		Code, found := response["code"]
		if !found {
			base.InfoLog(response)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 999,
				"msg":  "get user info failed",
				"data": nil,
			})
			c.Abort()
			return
		}
		code, ok := Code.(float64)
		if !ok {
			base.InfoLog(response)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  "system error",
				"data": nil,
			})
			c.Abort()
			return
		}
		if code > 0 {
			base.InfoLog(response)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 10000,
				"msg":  "token error",
				"data": nil,
			})
			c.Abort()
			return
		}

		Data, found := response["data"]
		if !found {
			base.InfoLog(response)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 999,
				"msg":  "get user info failed",
				"data": nil,
			})
			c.Abort()
			return
		}
		data, ok := Data.(map[string]interface{})
		if !ok {
			base.InfoLog(Data)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  "system error",
				"data": nil,
			})
			c.Abort()
			return
		}
		UserId, found := data["user_id"]
		if !found {
			base.InfoLog(response)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 999,
				"msg":  "get user info failed",
				"data": nil,
			})
			c.Abort()
			return
		}
		userId, ok := UserId.(string)
		if !ok {
			base.InfoLog(userId)
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 500,
				"msg":  "system error",
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Set("userId", userId)
	}
}
