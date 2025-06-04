package controller

import (
	"gin_test01/web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSeesion(ctx *gin.Context) {
	// 	初始化错误返回的map
	resp := make(map[string]string)
	// 调用utils包种的宏，go中称为常量
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}
