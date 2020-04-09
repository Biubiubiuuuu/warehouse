package userService

import (
	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/MD5Helper"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/jwtHelper"
	"github.com/Biubiubiuuuu/warehouse/server/models"
	"github.com/google/uuid"
)

// 登录
func LoginUser(login entity.UserLogin, ip string) (responseData entity.ResponseData) {
	if login.Tel == "" || login.Password == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	user := models.User{Tel: login.Tel, Password: MD5Helper.EncryptMD5To32Bit(login.Password)}
	if err := user.LoginUser(); err != nil {
		responseData.Message = msg.GetMsg(tcode.LOGIN_ERROR)
		return
	}
	token, err := jwtHelper.GenerateToken(login.Tel, login.Password)
	if err != nil {
		responseData.Message = msg.GetMsg(tcode.TOKEN_ERROR)
		return
	}
	// 写入uuid、token、IP，并返回用户信息
	uuid, _ := uuid.NewUUID()
	args := map[string]interface{}{"token": token, "ip": ip, "uuid": uuid}
	if err := user.UpdataUser(args); err != nil {
		responseData.Message = msg.GetMsg(tcode.LOGIN_ERROR)
		return
	}
	data := make(map[string]interface{})
	data["user"] = user
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.LOGIN_SUCCESS)
	responseData.Data = data
	return
}

// 注册
func RegisterUser(add entity.UserRegister) (responseData entity.ResponseData) {
	if add.Tel == "" || add.Password == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	newUser := models.User{Tel: add.Tel, Password: MD5Helper.EncryptMD5To32Bit(add.Password)}
	if err := newUser.RegisterUser(); err != nil {
		responseData.Message = msg.GetMsg(tcode.USERNAME_EXIST)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.ADD_SUCCESS)
	return
}

// 修改密码
func UpdateUserPass(updatePass entity.UserUpdatePass) (responseData entity.ResponseData) {
	if updatePass.Tel == "" || updatePass.NewPassword == "" || updatePass.OldPassword == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	user := models.User{Tel: updatePass.Tel, Password: MD5Helper.EncryptMD5To32Bit(updatePass.OldPassword)}
	if err := user.LoginUser(); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR)
		return
	}
	args := make(map[string]interface{})
	args["password"] = MD5Helper.EncryptMD5To32Bit(updatePass.NewPassword)
	if err := user.UpdataUser(args); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.UPDATE_SUCCESS)
	return
}
