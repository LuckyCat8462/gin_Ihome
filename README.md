# Home租房网(Rental Website)

## 项目简介

随着租房市场的日益发展，为满足用户便捷、高效的租房需求，本项目旨在构建一个功能完善、性能稳定的租房平台，提供房源展示、用户管理、订单处理等全方位服务。

## 系统架构

### 整体架构

![系统架构](https://example.com/architecture.png)<!-- 可根据实际情况添加架构图 -->

- **前端层**：使用Bootstrap框架构建响应式界面，提供用户友好的交互体验
- **Web服务层**：基于Gin框架实现的RESTful API服务，处理HTTP请求并调用微服务
- **微服务层**：采用Go-Micro框架实现的多个独立服务，通过gRPC通信
- **服务发现**：使用Consul实现服务注册与发现
- **数据存储层**：MySQL存储业务数据，Redis缓存热点数据，FastDFS存储图片文件

## 核心功能模块

### 1. 用户服务 (user)
- 用户注册、登录与退出
- 个人信息管理与修改
- 实名认证功能
- 头像上传与管理

### 2. 房源服务 (house)
- 房源信息发布与管理
- 房源图片上传（基于FastDFS）
- 房源搜索与筛选
- 房源详情展示
- 首页轮播图展示

### 3. 订单服务 (order)
- 房屋预订与订单创建
- 订单状态管理（接受/拒绝）
- 订单历史查询

### 4. 注册服务 (register)
- 用户注册流程处理
- 验证码验证

### 5. 验证码服务 (getCaptcha)
- 图片验证码生成与验证
- 短信验证码发送

## 项目结构

```
gin_Ihome/
├── conf/              # 配置文件
├── Learning/          # 学习示例代码
├── MyShell/           # 自动化部署脚本
├── service/           # 微服务模块
│   ├── getCaptcha/    # 验证码服务
│   ├── house/         # 房源服务
│   ├── order/         # 订单服务
│   ├── register/      # 注册服务
│   └── user/          # 用户服务
└── web/               # Web前端服务
    ├── conf/          # Web配置
    ├── controller/    # 控制器
    ├── model/         # 数据模型
    ├── proto/         # Protocol Buffers定义
    ├── utils/         # 工具函数
    ├── view/          # 前端视图文件
    └── main.go        # Web服务入口
```

## 技术栈

### 后端技术：
- **编程语言**：Go 1.24.4
- **Web框架**：Gin
- **微服务框架**：Go-Micro
- **服务发现**：Consul v1.21.1
- **ORM**：GORM
- **通信协议**：gRPC
- **数据序列化**：Protocol Buffers 21.2
- **数据库**：MySQL 8.0.42
- **缓存**：Redis 7.0.15 (go-redis v8.11.5, redigo v1.9.2)
- **文件存储**：FastDFS + Nginx 1.24.0

### 前端技术：
- **框架**：Bootstrap
- **语言**：HTML, CSS, JavaScript

## 安装与部署

### 环境要求

- Go 1.24.4或更高版本
- MySQL 8.0.42
- Redis 7.0.15
- Consul v1.21.1
- FastDFS + Nginx 1.24.0
- Protocol Buffers 21.2

### 安装步骤

1. **克隆项目代码**
   ```bash
   git clone https://github.com/LuckyCat8462/gin_Ihome.git
   cd gin_Ihome
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **配置数据库**
   - 创建MySQL数据库：`search_house`
   - 配置数据库连接信息（见web/conf/config.go）

4. **启动服务**
   ```bash
   # 启动微服务和相关组件
   cd MyShell
   ./ScriptHouses.sh
   
   # 启动Web服务
   cd ../web
   go run main.go
   ```

5. **访问系统**
   - 浏览器访问：http://localhost:8080

## API接口概览

### 用户相关接口
- `GET /api/v1.0/session` - 获取用户会话
- `POST /api/v1.0/users` - 用户注册
- `POST /api/v1.0/sessions` - 用户登录
- `DELETE /api/v1.0/session` - 用户登出
- `GET /api/v1.0/user` - 获取用户信息
- `PUT /api/v1.0/user/name` - 更新用户名
- `POST /api/v1.0/user/avatar` - 上传头像
- `POST /api/v1.0/user/auth` - 实名认证

### 房源相关接口
- `GET /api/v1.0/areas` - 获取地区列表
- `POST /api/v1.0/houses` - 发布房源
- `POST /api/v1.0/houses/:id/images` - 上传房源图片
- `GET /api/v1.0/houses/:id` - 获取房源详情
- `GET /api/v1.0/houses` - 搜索房源
- `GET /api/v1.0/house/index` - 获取首页轮播图

### 订单相关接口
- `POST /api/v1.0/orders` - 创建订单
- `GET /api/v1.0/user/orders` - 获取用户订单
- `PUT /api/v1.0/orders/:id/status` - 更新订单状态

### 验证码相关接口
- `GET /api/v1.0/imagecode/:uuid` - 获取图片验证码
- `GET /api/v1.0/smscode/:phone` - 获取短信验证码

## 系统特点

1. **微服务架构**：各模块独立部署，易于扩展和维护
2. **高性能**：使用Go语言开发，结合Redis缓存提升响应速度
3. **分布式文件存储**：FastDFS提供可靠的图片存储解决方案
4. **服务发现**：Consul实现服务的自动注册与发现
5. **自动化部署**：Shell脚本简化部署流程
6. **安全可靠**：完善的用户认证机制和数据验证

## 注意事项

1. 确保MySQL、Redis、Consul和FastDFS服务正常运行
2. 修改配置文件中的IP地址和端口以适配您的环境
3. 首次运行需要确保数据库已创建并配置正确的访问权限

## 文档说明

- API详细文档位于DOCS分支
- 完整的需求规格说明书、系统设计文档和用户手册将陆续更新

## 许可证

MIT License
