package base

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	jsoniter "github.com/json-iterator/go"
	"github.com/knight0zh/demo_config/config"
	"github.com/knight0zh/demo_pkg/pkg"
	"go.uber.org/zap"
)

const (
	YmdHis = "2006-01-02 15:04:05"
	Ymd    = "2006-01-02"
	His    = "15:04:05"
	YmdH00 = "2006-01-02 15:00:00"
	Ymd000 = "2006-01-02 00:00:00"
	H      = "15"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

type BaseContext struct {
	*gin.Context
}

func NewContext(ctx *gin.Context) *BaseContext {
	return &BaseContext{ctx}
}

func (uc *BaseContext) Success(obj interface{}) {
	rsp := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": obj,
	}
	uc.JSON(http.StatusOK, rsp)
	uc.Abort()
}

func (uc *BaseContext) Fail(code int, obj interface{}) {
	rsp := map[string]interface{}{
		"code": code,
		"msg":  pkg.Codes[code],
		"data": nil,
	}
	uc.JSON(http.StatusOK, rsp)
	uc.Abort()
}

func (uc *BaseContext) Alert(code int, msg string) {
	rsp := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": nil,
	}
	uc.JSON(http.StatusOK, rsp)
	uc.Abort()
}

func ErrLog(err error) {
	config.Logger.Error("ERROR", zap.Error(err))
}

func InfoLog(msg interface{}) {
	config.Logger.Info("INFO", zap.Reflect("msg", msg))
}

func init() {

	config.InitConfig(".")
	gin.SetMode(config.Get("app.mode").(string))

	config.InitOmsMysql()
	config.InitHdMysql()
	config.InitCommonRedis()
	binding.Validator = new(pkg.DefaultValidator)

	if err := config.InitAccessLogger(config.Log()); err != nil {
		log.Fatalf("init access logger failed, err:%v\n", err)
	}
	if err := config.InitErrorLogger(); err != nil {
		log.Fatalf("init sys logger failed, err:%v\n", err)
	}
}
