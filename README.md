基于Go的在线考试系统

1、如何启动
进入项目根目录 使用go mod init 初始化项目
使用go mod tidy下载所需库
配置config文件夹下application.yml文件 MySQL数据库配置文件
go mod ./main.go 运行

2、可以使用Docker打包项目
Dockerfile、Deployment.yarm已配置完毕

