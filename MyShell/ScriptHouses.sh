#! /bin/bash

#启动redis与consul服务发现
sudo redis-cli -h 192.168.81.128 -p 6349

sudo consul agnet-dev

#启动微服务

go run :~/Learning/WorkTools/Go_WorkSapce/src/gin_test01/service/getCaptcha/main.go

go run :~/Learning/WorkTools/Go_WorkSapce/src/gin_test01/service/register/main.go

go run :~/Learning/WorkTools/Go_WorkSapce/src/gin_test01/service/user/main.go

go run :~/Learning/WorkTools/Go_WorkSapce/src/gin_test01/service/house/main.go

go run :~/Learning/WorkTools/Go_WorkSapce/src/gin_test01/service/order/main.go
