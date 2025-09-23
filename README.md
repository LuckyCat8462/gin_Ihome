# Home租房网(Rental Website)

# 简述

​	这是一个使用Go语言构建的高性能租房网站服务。本项目提供完整的房源浏览、搜索、发布、用户管理、订单管理等功能。



# 功能特性

- 用户系统：用户注册、登录、实名认证、头像上传、个人信息管理
- 房源管理：发布、编辑、删除、上下架房源信息。支持使用FastDFS+Nginx实现多图上传。
- 智能搜索：支持按关键词、位置、价格范围、户型等条件筛选房源。
- 高性能：采用go-micro微服务框架，依托其高并发特性，轻松应对大量请求。



# 技术栈

## 后端：

- 语言：Go
- web框架：Gin
- 微服务框架：Go-Micro
- 服务发现：consul
- ORM:GORM
- 结构化数据传输:ProtoBuffer
- 远程调用协议：GRPC
- 数据库：MySQL
- 缓存：Redis
- 图片上传：fastdfs+nginx
- 脚本管理：Shell



# 安装与运行：

## 1、初始条件

- Go 1.24.4
- MySQL 8.0.42
- redis-cli 7.0.15
- Consul v1.21.1
- fastdfs-nginx-module-1.24
- nginx-1.24.0
- protobuf-all 21.2
- Grpc-Go V1.73
- go-redis v8.11.5
- redigo v1.9.2

## 2、获取代码

> git clone github.com/LuckyCat8462/gin_Ihome.git

## 3、安装依赖

在项目根目录src文件下输入go mod tidy,根据提示进行安装。

## 4、运行

（1）进入项目中的MyShell目录下，使用终端打开，并输入以下代码，启动redis、consul、以及项目相关的微服务等服务端所需内容。

 ./ScriptHouses.sh

(2)进入web目录下，使用go run main.go，运行项目客户端。

​	访问http://localhost:8080查看。



# 其他文档

接口文档位于DOCS分支。
需求规格说明书（SRS）、系统设计文档(SSD)、用户手册待更新。
