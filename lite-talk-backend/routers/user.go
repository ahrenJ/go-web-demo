package routers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lite-talk/models"
	"lite-talk/responseutil"
	"net/http"
	"strconv"
)

type User struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var body User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, responseutil.ErrorArgs())
		return
	}

	if models.QueryUserExisted(body.Nickname) {
		c.JSON(http.StatusOK, gin.H{"msg": "用户名已被注册"})
		return
	}

	encryptPwd, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	models.AddUser(body.Nickname, string(encryptPwd))
	c.JSON(http.StatusOK, responseutil.SuccessWithMsg("注册成功"))
}

func UserLogin(c *gin.Context) {
	var body User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, responseutil.ErrorArgs())
		return
	}

	userModel, existed := models.QueryUserByNickname(body.Nickname)
	if !existed {
		c.JSON(http.StatusBadRequest, responseutil.Error("用户名不存在"))
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userModel.EncryptPwd), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, responseutil.Error("密码错误"))
		return
	}
	c.JSON(http.StatusOK, responseutil.SuccessWithMsg("登录成功"))
}

func UserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responseutil.ErrorArgs())
		return
	}

	userModel, existed := models.QueryUserById(id)
	if !existed {
		c.JSON(http.StatusBadRequest, responseutil.Error("用户不存在"))
		return
	}

	var data = make(map[string]interface{})
	data["nickname"] = userModel.Nickname
	data["avatar"] = userModel.Avatar
	c.JSON(http.StatusOK, responseutil.SuccessWithData(data))
}
