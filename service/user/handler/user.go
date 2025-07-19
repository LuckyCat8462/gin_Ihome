package handler

import (
	"context"
	"fmt"
	"user/model"
	user "user/proto"
	"user/utils"
)

type User struct{}

// Return a new handler
func New() *User {
	return &User{}
}

func (e *User) Register(ctx context.Context, req *user.RegReq, rsp *user.Response) error {
	// 先校验短信验证码,是否正确. redis 中存储短信验证码.

	// 如果校验正确. 注册用户. 将数据写入到 MySQL数据库.

	err := model.RegisterUser(req.Mobile, req.Password)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
	} else {
		rsp.Errno = utils.RECODE_OK
		rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	}

	return nil
}

func (e *User) AuthUpdate(ctx context.Context, req *user.AuthReq, resp *user.AuthResp) error {

	//存储真实姓名和真是身份证号  数据库
	fmt.Println("req数据", req)

	err := model.SaveRealName(req.UserName, req.RealName, req.IdCard)

	if err != nil {
		fmt.Println("save错误", err)
		resp.Errno = utils.RECODE_DBERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	} else {
		fmt.Println("怎么的", err)
	}

	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	return nil
}
