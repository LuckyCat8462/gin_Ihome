# Home租房网(Rental Website)

# 简述

​	随着租房市场的日益发展，为满足用户便捷、高效的租房需求、打造一个功能完善、性能稳定的租房平台势在必行。本项目旨在构建一个集房源展示、用户管理、订单处理等多功能于一体的租房网站。	



# 项目内容

- **技术选型与架构搭建**：前端运用Bootstrap框架，为用户提供简洁美观、响应式的界面。后端采用Golang与Gin框架，结合RESTful API规范设计路由接口，确保接口的清晰与易用。使用Mysql数据库搭配GORM进行数据存储，保障数据的高效读写与管理。
- **微服务架构设计**：采用Go - Micro微服务架构，将系统拆分为用户模块、图片验证码获取模块、地区模块、房源模块、订单模块、设施模块等多个独立且松耦合的服务。通过Consul实现服务发现，使各个服务能够动态注册与发现，提高系统的可扩展性和容错性。利用grpc进行服务间的高效通信，确保数据传输的准确性和及时性。
- **文件存储与访问**：引入分布式文件系统FastDFS存储用户图片，并整合Nginx，实现通过HTTP协议快速访问存储在FastDFS中的文件，提升用户图片的加载速度。
- **自动化部署**：编写Shell脚本，实现微服务、Consul、FastDFS、Nginx的自动化启动，简化部署流程，提高开发和运维效率。



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
