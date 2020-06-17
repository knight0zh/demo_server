package base

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
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

var (
	Config      *config.ConfigBuild
	OmsMysql    *gorm.DB
	HdMysql     *gorm.DB
	CommonRedis *redis.Client
	Json        = jsoniter.ConfigCompatibleWithStandardLibrary
)

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
	pkg.Logger.Error("ERROR", zap.Error(err))
}

func InfoLog(msg interface{}) {
	pkg.Logger.Info("INFO", zap.Reflect("msg", msg))
}

func init() {

	Config = config.NewConfig(".")
	gin.SetMode(Config.Get("app.mode").(string))

	OmsMysql = Config.Mysql.OMS.NewMysql()
	HdMysql = Config.Mysql.OMS.NewMysql()
	CommonRedis = Config.Redis.Common.NewRedis()
	binding.Validator = new(pkg.DefaultValidator)

	if err := pkg.InitAccessLogger(Config.Log); err != nil {
		log.Fatalf("init access logger failed, err:%v\n", err)
	}
	if err := pkg.InitErrorLogger(); err != nil {
		log.Fatalf("init sys logger failed, err:%v\n", err)
	}
	watchConfig()
}

func watchConfig() {
	Config.Viper.WatchConfig()
	Config.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		OmsMysql.Close()
		HdMysql.Close()
		CommonRedis.Close()
		if err := pkg.InitAccessLogger(Config.Log); err != nil {
			log.Fatalf("init access logger failed, err:%v\n", err)
		}
		if err := pkg.InitErrorLogger(); err != nil {
			log.Fatalf("init sys logger failed, err:%v\n", err)
		}
	})
}
