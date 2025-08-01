package handler

import (
	"context"
	"user/model"
	user "user/proto"
	"user/utils"
)

type User struct{}

// Return a new handler
func New() *User {
	return &User{}
}

func (e *User) AuthUpdate(ctx context.Context, req *user.AuthReq, resp *user.AuthResp) error {
	//调用借口校验realName和idcard是否匹配

	//存储真实姓名和真是身份证号  数据库
	err := model.SaveRealName(req.UserName, req.RealName, req.IdCard)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	return nil
}
