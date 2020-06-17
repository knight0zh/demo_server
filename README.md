# go项目模版
## 介绍
基于gin框架简单封装的项目demo。该脚手架项目依赖demo-config和demo-pkg。

- 快速mysql table 转换 struct
- 支持配置热加载
- 日志分割
- token验证

## 快速开始
由于go modules 对私有仓库的限制。需要将进行以下转换。
```
go env -w GOPROXY=https://goproxy.cn,direct

cd demo-server
go mod tidy
go run main.go
```

## 快速mysql table转struct
脚手架集成快速转换struct工具提高开发效率。
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
