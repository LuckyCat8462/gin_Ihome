# 1、简介

​	这是Ihome租房网的API文档(V1.0)，用于项目后期明确请求参数及格式、梳理接口设计逻辑。

适用范围：项目全局

基础路径：http://localhost:8080/api/v1.0

最后更新：

维护人：neko

# 2、全局错误码

## 2.1.简述

​	此处规定了全局中可能出现的错误类型对应的错误码与英文。

| 错误码 | 错误解释               | 对应英文          |
| ------ | ---------------------- | ----------------- |
| 0      | 成功                   | RECODE_OK         |
| 4001   | 数据库查询错误         | RECODE_DBERR      |
| 4002   | 无数据                 | RECODE_NODATA     |
| 4003   | 数据错误               | RECODE_DATAERR    |
| 4004   | 数据错误               | RECODE_DATAERR    |
| 4101   | 用户未登录             | RECODE_SESSIONERR |
| 4102   | 用户登录失败           | RECODE_LOGINERR   |
| 4103   | 参数错误               | RECODE_PARAMERR   |
| 4104   | 用户不存在或未激活     | RECODE_USERONERR  |
| 4105   | 用户身份错误           | RECODE_ROLEERR    |
| 4106   | 密码错误               | RECODE_PWDERR     |
| 4107   | 用户已经注册           | RECODE_USERERR    |
| 4108   | 手机号错误             | RECODE_MOBILEERR  |
| 4201   | 非法请求或请求次数受限 | RECODE_REQERR     |
| 4202   | IP受限                 | RECODE_IPERR      |
| 4301   | 第三方系统错误         | RECODE_THIRDERR   |
| 4302   | 文件读写错误           | RECODE_IOERR      |
| 4500   | 内部错误               | RECODE_SERVERERR  |
| 4501   | 未知错误               | RECODE_UNKNOWERR  |



# 3、首页模块

## 3.1、web服务

**简要描述：**

- 网站web页面显示与路由解析

**请求URL：** 

- `http://localhost:8080`

**路由列表：**

| **服务编号** | **服务名称**                 | **请求类型** | **url**                     | **调用函数**    |
| ------------ | ---------------------------- | ------------ | --------------------------- | --------------- |
| 01           | web服务                      |              | /                           |                 |
| 02           | 获取地区信息服务             | GET          | api/v1.0/areas              | GetArea         |
| 03           | 获取验证码图片服务           | GET          | api/v1.0/imagecode/:uuid    | GetImageCd      |
| 04           | 获取短信验证码服务           | GET          | api/v1.0/smscode/:mobile    | GetSmscd        |
| 05           | 发送注册信息服务             | POST         | api/v1.0/users              | PostRet         |
| 06           | 获取session信息服务          | GET          | api/v1.0/session            | GetSession      |
| 07           | 发送登陆信息服务             | POST         | api/v1.0/sessions           | PostLogin       |
| 08           | 删除（退出）登陆信息服务     | DELETE       | api/v1.0/session            | DeleteSession   |
| 09           | 获取用户基本信息服务         | GET          | api/v1.0/user               | GetUserInfo     |
| 10           | 发送（上传）用户头像服务     | POST         | api/v1.0/user/avatar        | PostAvatar      |
| 08           | 更新用户名服务               | PUT          | api/v1.0/user/name          | PutUserInfo     |
| 12           | 获取（检查）用户实名信息服务 | GET          | api/v1.0/user/auth          | GetUserAuth     |
| 13           | 更新用户实名认证信息服务     | PUT          | api/v1.0/user/auth          | PutUserAuth     |
| 14           | 获取用户已发布房源信息服务   | GET          | api/v1.0/user/houses        | GetUserHouses   |
| 15           | 发送（发布）房源信息服务     | POST         | api/v1.0/houses             | PostHouses      |
| 16           | 发送（上传）房屋图片服务     | POST         | api/v1.0/houses/:id/images  | PostHousesImage |
| 17           | 获取房屋详细信息服务         | GET          | api/v1.0/houses/:id         | GetHouseInfo    |
| 18           | 获取首页轮播图片服务         | GET          | api/v1.0/house/index        | GetIndex        |
| 19           | 获取（搜索）房源服务         | GET          | api/v1.0/houses             | GetHouses       |
| 20           | 发送（发布）订单服务         | POST         | api/v1.0/orders             | PostOrders      |
| 21           | 获取房东/租户订单信息服务    | GET          | api/v1.0/user/orders        | GetUserOrder    |
| 22           | 更新房东同意/拒绝订单        | PUT          | api/v1.0/orders/:id/status  | PutOrders       |
| 23           | 更新用户评价订单信息         | PUT          | api/v1.0/orders/:id/comment | PutComment      |

- 备注：分清请求类型与参数

## 3.2、获取地区信息服务

**简要描述：**

- 获取相关地域信息

**请求URL：** 

- `http://localhost/api/v1.0/areas`

**请求方式：**

- GET 

**参数：**

无

**返回成功**

```
{ 	"errno": 0,

	"errmsg":"OK",

	"data": [   {"aid": 1, "aname": "东城区"}, 

	{"aid": 2, "aname": "西城区"},

	{"aid": 3, "aname": "通州区"},

	{"aid": 4, "aname": "顺义区"}] 

}
```

**返回失败**

```
{   "errno": "4001",   //状态码  
"errmsg":"状态错误信息" //状态信息 
}
```

**返回参数说明**

| **参数名** | **类型**     | **说明** |
| ---------- | ------------ | -------- |
| errno      | string       | 状态码   |
| errmsg     | string       | 状态信息 |
| data       | 切片         | 地域信息 |
| aid        | int32（int） | 地域编号 |
| aname      | string       | 地域名   |

**备注**

- 返回给前端与proto是不一样的

## 3.3、获取首页轮播图片服务

**简要描述：**

- 获取首页轮播图以及相关房屋图片

**请求URL：** 

- `http://localhost/api/v1.0/house/index`

**请求方式：**

- GET 

**参数：**

无

**返回成功**

```
{
    "errno": "0",
    "errmsg": "成功",
    "data": {
        "houses": [
            {
                "house_id": this.Id,
                "title": this.Title,
                "price": this.Price,
                "area_name": this.Area.Name,
                "img_url": utils.AddDomain2Url(this.Index_image_url),
                "room_count": this.Room_count,
                "order_count": this.Order_count,
                "address": this.Address,
                "user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
                "ctime": this.Ctime.Format("2006-01-02 15:04:05"),
                
            },
            {
                "house_id": this.Id,
                "title": this.Title,
                "price": this.Price,
                "area_name": this.Area.Name,
                "img_url": utils.AddDomain2Url(this.Index_image_url),
                "room_count": this.Room_count,
                "order_count": this.Order_count,
                "address": this.Address,
                "user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
                "ctime": this.Ctime.Format("2006-01-02 15:04:05"),
                
            }
        ],
        
    }
}
```

**返回失败**

```
{
    "errno": "400x",     //状态码
    "errmsg":"状态错误信息"  //状态信息
}
```

# 4、注册模块

## 4.1、获取验证码图片服务

**简要描述：**

- 通过调用库完成随机数与图片验证码的生成

**请求URL：** 

- `http://localhost/api/v1.0/imagecode/:uuid`

**请求方式：**

- GET 

**参数：**

Uuid

**返回成功**

| {<br/>    "errno": "0",   //状态码<br/>    "errmsg":"成功"<br/>    二进制图片数据<br/>} |
| ------------------------------------------------------------ |
| ` `                                                          |

**返回失败**

```
{
    "errno": "400x",   //状态码
    "errmsg":"状态错误信息"
}
```

**返回参数说明**

无

**备注**

- 成功返回为纯二进制图片数据



# 5、登录模块

## 5.1、获取session信息服务

**简要描述：**

- 获取用户登录的session信息

**请求URL：** 

- `http://localhost/api/v1.0/session`

**请求方式：**

- GET 

**参数：**

无

**返回成功**

```
{   "errno": "0",  
"errmsg":"OK", 
"data": {"name" : "13313331333"} }
```

**返回失败**

```
{   "errno": "400x",  //状态码
"errmsg":"状态错误信息" }
```

## 5.2、发送登录信息服务

**简要描述：**

- 发送用户登录信息进行登录

**请求URL：** 

- `http://localhost/api/v1.0/sessions`

**请求方式：**

- POST 

**参数：**

```
{   mobile: "133",    password: "itcast" }
```

 

| **参数名** | **必选** | **类型** | **说明** |
| ---------- | -------- | -------- | -------- |
| mobile     | 是       | string   | 手机号   |
| password   | 是       | string   | 密码     |

**返回成功**

```
{    "errno": "0",   "errmsg":"OK", }
```

**返回失败**

```
{   "errno": "400x",  //状态码   "errmsg":"状态错误信息" }
```

## 5.3、退出登录信息服务

**简要描述：**

- 删除用户登录的session信息

**请求URL：** 

- `http://localhost/api/v1.0/session`

**请求方式：**

- DELETE

**参数：**

无

**返回成功**

```
{   "errno": "0",   "errmsg":"OK", }
```

**返回失败**

```
{    "errno": "400x",  //状态码   "errmsg":"状态错误信息" }
```

**备注**

- 请求成功后自动退出登录

# 6、用户模块

## 6.1、获取用户基本信息服务

**简要描述：**

- 获取用户基本信息进行显示

**请求URL：** 

- `http://localhost/api/v1.0/user`

**请求方式：**

- GET

**参数：**

无

**返回成功**

```
{
  "errno": "0",
  "errmsg": "成功",
  "data": {
    "user_id": 1,
    "name": " 猫猫",
    "mobile": "080",
    "real_name": "熊猫",
    "id_card": "210082244556677",
    "avatar_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1n7It2ANn1dAADexS5wJKs808.png"
  }
}
```

**返回失败**

```
{   "errno": "400x",  //状态码  
"errmsg":"状态错误信息" }
```

## 6.2、发送（上传）用户头像服务

**简要描述：**

- 调用fastdfs接口上传用户头像

**请求URL：** 

- `http://localhost/api/v1.0/user/avatar`

**请求方式：**

- POST 

**参数：**

二进制图片数据

**返回成功**

```
{
  "errno": "0",
  "errmsg": "成功",
  "data": {
    "avatar_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1n6_L-AOB04AADexS5wJKs662.png" //图片地址需要进行拼接
  } 
}
```

**返回失败**

```
{   "errno": "400x",  //状态码  
"errmsg":"状态错误信息" }
```

## 6.3、更新用户名服务

**简要描述：**

- 将用户名进行更新

**请求URL：** 

- `http://localhost/api/v1.0/user/name`

**请求方式：**

- PUT

**参数：**

```
{  "name":" 猫猫" }
```

 

| **参数名** | **必选** | **类型** | **说明** |
| ---------- | -------- | -------- | -------- |
| name       | 是       | string   | 用户名   |

**返回成功**

```
{  "errno": "0",  "errmsg": "成功",  "data": {   "name": " 猫猫"  } }
```

**返回失败**

```
{   "errno": "400x",  //状态码   "errmsg":"状态错误信息" }
```

# 7、认证相关

## 7.1、获取（检查）用户实名信息服务

**简要描述：**

- 获取用户基本信息进行检查

**请求URL：** 

- `http://localhost/api/v1.0/user/auth`

**请求方式：**

- GET

**参数：**

无

**返回成功**

```
{
  "errno": "0",
  "errmsg": "成功",
  "data": {
    "user_id": 1,
    "name": "猫猫",
    "mobile": "13086162323",
    "real_name": "猫猫",
    "id_card": "123456789123456789",
    "avatar_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1n7It2ANn1dAADexS5wJKs808.png"
  }
}
```

**返回失败**

```
{   "errno": "400x",  //状态码   "errmsg":"状态错误信息" }
```

## 7.2、发送用户实名认证信息服务

**简要描述：**

- 进行实名认证操作

**请求URL：** 

- `http://localhost/api/v1.0/user/auth`

**请求方式：**

- POST 

**参数：**

```
{   real_name: "猫猫   
id_card: "123456789123456789" }
```

 

| **参数名** | **必选** | **类型** | **说明**   |
| ---------- | -------- | -------- | ---------- |
| real_name  | 是       | string   | 真实姓名   |
| id_card    | 是       | string   | 身份证号码 |

**返回成功**

```
{  "errno": "0", 
"errmsg": "成功" }
```

**返回失败**

```
{   "errno": "400x",  //状态码
"errmsg":"状态错误信息" }
```

# 8、商品相关

## 8.1、获取用户已发布房源信息服务

**简要描述：**

- 用户注册接口

**请求URL：** 

- `http://localhost/api/v1.0/user/houses`

**请求方式：**

- GET

**参数：**

无

**返回示例**

```
{
  "errno": "0",
  "errmsg": "成功",
  "data": {
    "houses": [
      {
        "address": "上海市",
        "area_name": "松江区",
        "ctime": "2025-8-12 21:16:24",
        "house_id": 1,
        "img_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBJY-AL3m8AAS8K2x8TDE052.jpg",
        "order_count": 0,
        "price": 399,
        "room_count": 2,
        "title": "开元地中海",
        "user_avatar": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBLFeALIEjAADexS5wJKs340.png"
      }
    ]
  }
}
```

## 8.2、发送（发布）房源信息服务

**简要描述：**

- 将用户输入的信息发送到后台进行保存

**请求URL：** 

- `http://localhost/api/v1.0/houses`

**请求方式：**

- POST 

**参数：**

```
{
"title":"开元地中海",
"price":"339",
"area_id":"13",
"address":"松江区北翠路1000弄",
"room_count":"2",
"acreage":"60",
"unit":"2室1厅",
"capacity":"3",
"beds":"双人床2张",
"deposit":"200",
"min_days":"3",
"max_days":"0",
"facility":["1","2","3","7","12","14","16","17","18","21","22"]
}
```

**返回成功**

```
{  "errno": "0",  
"errmsg": "成功"  
"data" :{      "house_id": "1"  } }
```

**返回失败**

```
{   "errno": "400x",  //状态码   
"errmsg":"状态错误信息" }
```

## 8.3、发送（上传）房屋图片服务

**简要描述：**

- 上传房屋图片

**请求URL：** 

- `http://localhost/api/v1.0/houses/:id/images`
- `http://localhost/api/v1.0/houses/3/images`
- `3表示房源id`

**请求方式：**

- POST 

**参数：** 二进制图片数据

**返回成功**

```
{  "errno": "0", 
"errmsg": "成功", 
"data": {   "url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBLmWAHlsrAAaInSze-cQ719.jpg"  } }
```

**返回失败**

```
{  "errno": "0",  "errmsg": "成功",  "data": {   "url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBLmWAHlsrAAaInSze-cQ719.jpg"  } }
```

## 8.4、获取房屋详细信息服务

**简要描述：**

- 将对应编号的房屋信息获取到后发送给浏览器

**请求URL：** 

- `http://localhost/api/v1.0/houses/:id`
- `http://localhost/api/v1.0/houses/1`
- `1表示房源id`

**请求方式：**

- GET

**参数：**

无

**返回成功**

```
{
  "errno": "0",
  "errmsg": "成功",
  "data": {
    "house": {
      "acreage": 80,
      "address": "上海市松江区",
      "beds": "2双人床",
      "capacity": 3,
      "comments": [
        {
          "comment": "评论的内容",
          "ctime": "2025-08-12 12:30:30",
          "user_name": "zhi"
        },
        {
          "comment": "评论的内容",
          "ctime": "2025-08-12 12:30:30",
          "user_name": "zhi"
        }
      ],
      "deposit": 200,
      "facilities": [9,08,13,16,19,20,21,23],
      "hid": 1,
      "img_urls": [
        "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBJY-AL3m8AAS8K2x8TDE052.jpg",
        "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBJZmAYqGWAAaInSze-cQ230.jpg"
      ],
      "max_days": 30,
      "min_days": 1,
      "price": 100,
      "room_count": 2,
      "title": "开元地中海",
      "unit": "3室3厅",
      "user_avatar": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBLFeALIEjAADexS5wJKs340.png",
      "user_id": 1,
      "user_name": "Panda"
    },
    "user_id": 1
  }
}

```

**返回失败**

```
{   "errno": "400x",  //状态码   "errmsg":"状态错误信息" }
```

## 8.5、获取（搜索）房源服务

**简要描述：**

- 根据条件搜索房源信息

**请求URL：** 

- `http://localhost/api/v1.0/houses`
- `http://localhost/api/v1.0/houses?aid=5&sd=2025-08-12&ed=2025-08-30&sk=new&p=1`
- `adi表示地区编号`
- `sd表示起始日期`
- `ed表示结束日期`
- `sk表示查询方式`
- `p表示页码`

**请求方式：**

- GET

**参数：**

无

**返回示例**

```
{  "errno": "0", 
	"errmsg": "成功", 
	"data": {   "current_page": 1,  
	"houses": [  
{     "address": "上海市",  
    "area_name": "松江区",    
    "ctime": "2025-08-13 08:16:24",   
    "house_id": 1,   
    "img_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBJY-AL3m8AAS8K2x8TDE052.jpg",   
    "order_count": 0,    
    "price": 100,   
    "room_count": 2,   
    "title": "开元地中海13号楼",   
    "user_avatar": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBLFeALIEjAADexS5wJKs340.png"    },
}
```

**返回参数说明**

| **参数名** | **类型** | **说明** |
| ---------- | -------- | -------- |
| total_page | int      | 总页数   |

# 9、订单相关

## 9.1、发送（发布）订单服务

**简要描述：**

- 对商品进行下单操作

**请求URL：** 

- `http://localhost/api/v1.0/orders`

**请求方式：**

- POST 

**参数：**

```
{  "house_id": "1", 
"start_date": "2025-08-08 21:23:49", 
"end_date": "2025-08-12 21:23:49", }
```

 

| **参数名** | **必选** | **类型** | **说明** |
| ---------- | -------- | -------- | -------- |
| house_id   | 是       | string   | 商品id   |
| start_date | 是       | string   | 起始时间 |
| end_date   | 是       | string   | 结束时间 |

**返回成功**

```
{  "errno": "0",  "errmsg": "成功",  "data": {   "order_id":"1"  } }
```

**返回失败**

```
{   "errno": "400x",  //状态码   "errmsg":"状态错误信息" }
```

## 9.2、获取房东/租户订单信息服务

**简要描述：**

- 区分角色查看订单信息

**请求URL：** 

- `http://localhost/api/v1.0/user/orders?role=custom`
- `http://localhost/api/v1.0/user/orders?role=landlord`
- `custom为租客查看订单信息`
- `landlord为房东查看订单信息`

**请求方式：**

- GET 

**参数：**

无

**返回成功**

```
{
  "errno": "0",
  "errmsg": "成功",
  "data": {
    "orders": [
      {
        "amount": 200,
        "comment": "哈哈拒接",
        "ctime": "2025-08-08 21:23:49",
        "days": 2,
        "end_date": "2025-08-29 16:00:00",
        "img_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBJY-AL3m8AAS8K2x8TDE052.jpg",
        "order_id": 3,
        "start_date": "2025-08-28 16:00:00",
        "status": "REJECTED",//WAIT_ACCPET,WAIT_COMMENT,REJECTED,COMPLETE,CANCELED
        "title": "开元地中海"
      },
      {
        "amount": 1500,
        "comment": "",
        "ctime": "2025-08-08 01:32:10",
        "days": 15,
        "end_date": "2025-08-24 16:00:00",
        "img_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBJY-AL3m8AAS8K2x8TDE052.jpg",
        "order_id": 2,
        "start_date": "2025-08-10 16:00:00",
        "status": "WAIT_COMMENT",
        "title": "开元地中海"
      },
      {
        "amount": 300,
        "comment": "",
        "ctime": "2025-08-10 01:46:00",
        "days": 3,
        "end_date": "2025-08-08 16:00:00",
        "img_url": "http://192.168.81.128:9998/group1/M00/00/00/Zciqq1oBJY-AL3m8AAS8K2x8TDE052.jpg",
        "order_id": 1,
        "start_date": "2025-08-09 16:00:00",
        "status": "WAIT_COMMENT",
        "title": "开元地中海"
      }
    ]
  }
}
	
```

**返回失败**

```
{   "errno": "400x",  //状态码  
"errmsg":"状态错误信息" }
```

## 9.3、更新房东同意/拒绝订单

**简要描述：**

- 用户注册接口

**请求URL：** 

- `http://localhost/api/v1.0/orders/4/status`
- `http://localhost/api/v1.0/orders/:id/status`
- `4表示订单id`

**请求方式：**

- PUT

**参数：**

```
{action: "accept"} 
```

 

| **参数名** | **必选** | **类型** | **说明**                           |
| ---------- | -------- | -------- | ---------------------------------- |
| action     | 是       | string   | "accept"表示接受, "reject"表示拒绝 |

**返回成功**

```
{  "errno": "0", 
"errmsg": "成功" }
```

**返回失败**

```
{   "errno": "400x",  //状态码 
"errmsg":"状态错误信息" }
```

## 9.4、更新用户评价订单信息

**简要描述：**

- 对已完成订单进行评价

**请求URL：** 

- `http://localhost/api/v1.0/orders/:id/comment`
- `http://localhost/api/v1.0/orders/2/comment`
- `2表示订单id`

**请求方式：**

- PUT 

**参数：**

{ order_id: "2", comment: "烂房子！" }

| **参数名** | **必选** | **类型** | **说明** |
| ---------- | -------- | -------- | -------- |
| order_id   | 是       | string   | 订单编号 |
| comment    | 是       | string   | 评论内容 |

**返回成功**

```
{  "errno": "0", 
"errmsg": "成功" }
```

**返回失败**

```
{    "errno": "400x",  //状态码 
"errmsg":"状态错误信息" }
```

 
