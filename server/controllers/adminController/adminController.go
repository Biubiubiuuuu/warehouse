package adminController

import (
	"net/http"

	"github.com/Biubiubiuuuu/warehouse/server/helpers/jwtHelper"
	"github.com/google/uuid"

	"github.com/Biubiubiuuuu/warehouse/server/common/response"
	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/MD5Helper"
	"github.com/Biubiubiuuuu/warehouse/server/models"
	"github.com/gin-gonic/gin"
)

// 登录请求结构体
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 注册请求结构体
type Register struct {
	OnlineUsername string `json:"online_username"` // 当前登录用户名
	Username       string `json:"username"`        // 用户名
	Password       string `json:"password"`        // 密码
	Administrator  string `json:"administrator"`   // 超级管理员 Y | N
}

// @Summary 管理员登录
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param body body adminController.Login true "body"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS"} {"status": false, "message": "用户名或密码不能为空"}  {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/Login [POST]
// @Security ApiKeyAuth
func LoginAdmin(c *gin.Context) {
	status := false
	code := tcode.SUCCESS
	data := make(map[string]interface{})
	req := &Login{}
	// 参数校验
	if c.ShouldBindJSON(&req) != nil {
		code = tcode.PARAMETER_ERROR
	} else if req.Username == "" || req.Password == "" {
		code = tcode.NOTNULL
	}
	if code == tcode.SUCCESS {
		admin := models.Admin{Username: req.Username, Password: MD5Helper.EncryptMD5To32Bit(req.Password)}
		if err := admin.Login(); err == nil {
			if token, err := jwtHelper.GenerateToken(req.Username, req.Password); err == nil {
				// 写入uuid、token、IP，并返回用户信息
				uuid, _ := uuid.NewUUID()
				args := map[string]interface{}{"token": token, "ip": c.ClientIP(), "uuid": uuid}
				if err := admin.UpdataInfo(args); err == nil {
					data["user"] = admin
					status = true
				} else {
					code = tcode.ERROR
				}
			} else {
				code = tcode.TOKEN_ERROR
			}
		} else {
			code = tcode.ERROR
		}
	}
	message := msg.GetMsg(code)
	responseJson := response.ResponseJson(status, data, message)
	c.JSON(http.StatusOK, responseJson)
}

// @Summary 添加管理员
// @tags 管理员
// @Accept  application/json
// @Produce  json
// @Param body body adminController.Register true "body"
// @Success 200 {string} json "{"status": true, "message": "SUCCESS"} {"status": false, "message": "用户名或密码不能为空"}  {"status": false, "message": "ERROR"}"
// @Router /api/v1/admin/addAdmin [POST]
// @Security ApiKeyAuth
func AddAdmin(c *gin.Context) {
	code := tcode.SUCCESS
	message := msg.GetMsg(code)
	responseJson := response.ResponseJson(true, nil, message)
	req := &Register{}
	// 参数校验
	if c.ShouldBindJSON(&req) != nil {
		code = tcode.PARAMETER_ERROR
	} else if req.Username == "" || req.Password == "" {
		code = tcode.NOTNULL
	}
	if code == tcode.SUCCESS {
		admin := models.Admin{Username: req.OnlineUsername}
		if admin.CheckAdministrator() {
			if req.Administrator != "Y" {
				req.Administrator = "N"
			}
			admin = models.Admin{Username: req.Username, Password: MD5Helper.EncryptMD5To32Bit(req.Password), Administrator: req.Administrator}
			if err := admin.Register(); err == nil {
				c.JSON(http.StatusOK, responseJson)
			}
			code = tcode.ERROR
		}
		code = tcode.NOTADMINISTRATOR
	}
	message = msg.GetMsg(code)
	responseJson = response.ResponseJson(false, nil, message)
	c.JSON(http.StatusOK, responseJson)
}
