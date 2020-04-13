package adminService

import (
	tcode "github.com/Biubiubiuuuu/warehouse/server/common/tips/code"
	"github.com/Biubiubiuuuu/warehouse/server/common/tips/msg"
	"github.com/Biubiubiuuuu/warehouse/server/entity"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/MD5Helper"
	"github.com/Biubiubiuuuu/warehouse/server/helpers/jwtHelper"
	"github.com/Biubiubiuuuu/warehouse/server/models"
	"github.com/google/uuid"
)

// 管理员登录
func LoginAdmin(login entity.Login, ip string) (responseData entity.ResponseData) {
	if login.Username == "" || login.Password == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	admin := models.Admin{Username: login.Username, Password: MD5Helper.EncryptMD5To32Bit(login.Password)}
	if err := admin.LoginAdmin(); err != nil {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_ERROR)
		return
	}
	token, err := jwtHelper.GenerateToken(login.Username, login.Password)
	if err != nil {
		responseData.Message = msg.GetMsg(tcode.TOKEN_ERROR)
		return
	}
	// 写入uuid、token、IP，并返回用户信息
	uuid, _ := uuid.NewUUID()
	args := map[string]interface{}{"token": token, "ip": ip, "uuid": uuid}
	if err := admin.UpdataAdminInfo(args); err != nil {
		responseData.Message = msg.GetMsg(tcode.LOGIN_ERROR)
		return
	}
	data := make(map[string]interface{})
	data["user"] = admin
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.LOGIN_SUCCESS)
	responseData.Data = data
	return
}

// 添加管理员账号
func AddAdmin(add entity.Register) (responseData entity.ResponseData) {
	if add.Username == "" || add.Password == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	if add.Administrator != "Y" {
		add.Administrator = "N"
	}
	newAdmin := models.Admin{Username: add.Username, Password: MD5Helper.EncryptMD5To32Bit(add.Password), Administrator: add.Administrator}
	if err := newAdmin.RegisterAdmin(); err != nil {
		responseData.Message = msg.GetMsg(tcode.USERNAME_EXIST)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.ADD_SUCCESS)
	return
}

// 查询管理员账号
func QueryByLimitOffset(pageSize int, page int) (responseData entity.ResponseData) {
	admins := models.QueryAdminByLimitOffset(pageSize, page)
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	if len(admins) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	count := models.QueryAdminCount()
	data := make(map[string]interface{})
	data["users"] = admins
	data["count"] = count
	responseData.Data = data
	responseData.Status = true
	return
}

// 修改管理员密码
func UpdateAdminPass(token string, updatePass entity.UpdatePass) (responseData entity.ResponseData) {
	if updatePass.NewPassword == "" || updatePass.OldPassword == "" {
		responseData.Message = "密码不能为空！"
		return
	}
	admin := models.Admin{Token: token}
	if !admin.QueryAdminByUsername() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + "，管理员不存在"
		return
	}
	admin = models.Admin{Username: admin.Username, Password: MD5Helper.EncryptMD5To32Bit(updatePass.OldPassword)}
	if err := admin.LoginAdmin(); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR) + "；" + msg.GetMsg(tcode.OLD_ERROR)
		return
	}
	args := make(map[string]interface{})
	args["password"] = MD5Helper.EncryptMD5To32Bit(updatePass.NewPassword)
	if err := admin.UpdataAdminInfo(args); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.UPDATE_SUCCESS)
	return
}

// 删除
func DeleteAdmin(ids []int64) (responseData entity.ResponseData) {
	if len(ids) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTNULL)
		return
	}
	admin := models.Admin{}
	if err := admin.DeleteAdmins(ids); err != nil {
		responseData.Message = msg.GetMsg(tcode.DELETE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.DELETE_SUCCESS)
	return
}
