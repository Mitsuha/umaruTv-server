package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"umarutv/app/common/Forms"
	"umarutv/common/auth"
	"umarutv/common/dao"
	common "umarutv/common/gin"
)

func HttpLogin(c *gin.Context) {
	var form Forms.LoginForm

	_ = c.ShouldBind(&form)
	if err := form.Check(); err != nil {
		common.JSONResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := dao.FullUserByName(form.Username)
	if err != nil {
		common.JSONResponse(c, http.StatusForbidden, err.Error())
		return
	}
	token, err := auth.AttemptLogin(user, form.Password)
	if err != nil {
		common.JSONResponse(c, http.StatusForbidden, err.Error())
		return
	}

	login := ResponseAfterLogin{}
	common.JSONResponse(c, http.StatusOK, login.LoadFullUser(user, token))
}

func HttpRegister(c *gin.Context) {
	var form RegisterForm
	if err := c.ShouldBind(&form); err != nil {
		common.JSONResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := form.Check(); err != nil {
		common.JSONResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user := form.GenerateAUser()
	user.Password = auth.EncryptionPassword(form.Password)

	if err := dao.CreateUserFromFullUser(user); err != nil {
		common.JSONResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, _ := auth.AttemptLogin(user, form.Password)

	login := ResponseAfterLogin{}
	common.JSONResponse(c, http.StatusOK, login.LoadFullUser(user, token))
}
