package controller

import (
	"cloud/common/utils/ctx"
	"cloud/error_code"
	"cloud/module/user_manage/pojo"
	"cloud/module/user_manage/service"
	"github.com/gin-gonic/gin"
)

// Login One godoc
// @Summary      用户登录
// @Description  用户登录
// @Tags         user_manage
// @Accept       json
// @Produce      plain
// @Param        req    body      pojo.LoginReq  true  "LoginReq"
// @Success      200  {object}  pojo.ResponseTemplate “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /user/login [post]
func Login(c *gin.Context) {
	req := pojo.LoginReq{}
	//参数校验解析
	if err := c.ShouldBind(&req); nil != err {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}
	code, result := service.Login(req)

	ctx.ResponseJson(c, code, result)
}

// Register One godoc
// @Summary      用户注册
// @Description  用户注册
// @Tags         user_manage
// @Accept       json
// @Produce      plain
// @Param        req    body      pojo.RegisterReq  true  "RegisterReq"
// @Success      200  {object}  pojo.ResponseTemplate “successful”
// @Failure      400  {object}  pojo.ResponseTemplate
// @Router       /user/register [post]
func Register(c *gin.Context) {
	req := pojo.RegisterReq{}
	if err := c.ShouldBind(&req); err != nil {
		ctx.ResponseJson(c, error_code.ErrInvalidJsonFormat, nil)
		return
	}
	//处理用户注册
	code := service.Register(req)
	ctx.ResponseJson(c, code, nil)
}
