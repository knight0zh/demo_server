# go项目模版
## 介绍
基于gin框架简单封装的项目demo。该脚手架项目依赖 https://github.com/knight0zh/demo_config 和 https://github.com/knight0zh/demo_pkg

- 快速`mysql table` 转换 `struct`
- 配置热加载(后期支持配置中心热加载)
- 日志分割
- 分库分表逻辑
- `token`验证

## 目录结构
    demo-server
    ├── Dockerfile  // 生成docker镜像
    ├── Makefile    // 编译构建
    ├── README.md  
    ├── base        // 一些基础初始化
    │   └── base.go
    ├── config.yml  // 配置文件
    ├── demo_server // bin文件
    ├── gen         // mysqsl table转struct工具
    ├── go.mod
    ├── go.sum
    ├── install.sh  // 编译脚本
    ├── logs        // 日志
    │   ├── access-20200617.log
    │   ├── access.log -> logs/access-20200617.log
    │   ├── error-20200617.log
    │   └── error.log -> logs/error-20200617.log
    ├── main.go
    ├── middlewares // 中间件
    │   ├── auth.go
    │   └── zlog.go
    ├── models     
    │   ├── model_base.go
    │   └── tb_csc_online_repair.go
    ├── routers     // 路由
    │   ├── api     // 存放具体controller
    │   │   └── demo
    │   │       └── demo.go
    │   └── router.go
    ├── service     // 处理具体业务逻辑
    │   └── demo.go
    └── template    // gen模版
        ├── mapping.json
        └── model.go.tmpl


## 快速开始
```
go env -w GOPROXY=https://goproxy.cn,direct

cd go-demo-server
go mod tidy
go run main.go
```

## 快速mysql table转struct
脚手架集成快速转换`struct`工具提高开发效率。
```
cd go-demo-server
./gen  --connstr "user:pwd@tcp(host:port)/dbname?parseTime=True" --json --templateDir=template --overwrite --out=. --model=models --mapping=template/mapping.json --database dbname -t tbname
```

## 快速编译
```
make install_all
```

## Docker构建
```
make docker_build
docker run -p 8088:8088 go-server 
```

## Supervisor配置
```
user=
command= /.../your-project/{bin所在目录}
directory=/.../your-project/{bin所在目录}                                                                                                                              
autostart=true
autorestart=true
redirect_stderr=True
stdout_logfile=/.../{脚本名称}/{日志}.log
```