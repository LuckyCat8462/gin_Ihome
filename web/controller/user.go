package controller

import (
	"fmt"
	"gin_test01/web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取session信息
func GetSession(ctx *gin.Context) {
	// 	初始化错误返回的map
	resp := make(map[string]string)
	// 调用utils包种的宏，go中称为常量
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}

// 获取图片信息
func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	fmt.Println("uuid", uuid)
}
