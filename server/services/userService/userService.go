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
func LoginUser(request entity.UserLogin, ip string) (responseData entity.ResponseData) {
	if request.Tel == "" || request.Password == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	user := models.User{Tel: request.Tel, Password: MD5Helper.EncryptMD5To32Bit(request.Password)}
	if err := user.LoginUser(); err != nil {
		responseData.Message = msg.GetMsg(tcode.LOGIN_ERROR)
		return
	}
	token, err := jwtHelper.GenerateToken(request.Tel, request.Password)
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
func RegisterUser(request entity.UserRegister) (responseData entity.ResponseData) {
	if request.Tel == "" || request.Password == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	newUser := models.User{Tel: request.Tel, Password: MD5Helper.EncryptMD5To32Bit(request.Password)}
	if err := newUser.RegisterUser(); err != nil {
		responseData.Message = msg.GetMsg(tcode.USERNAME_EXIST)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.ADD_SUCCESS)
	return
}

// 修改密码
func UpdateUserPass(token string, request entity.UserUpdatePass) (responseData entity.ResponseData) {
	if request.NewPassword == "" || request.OldPassword == "" {
		responseData.Message = msg.GetMsg(tcode.USERNAME_OR_PASSWORD_NOT_NULL)
		return
	}
	user := models.User{Token: token}
	if !user.QueryUser() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + "，用户不存在"
		return
	}
	user = models.User{Tel: user.Tel, Password: MD5Helper.EncryptMD5To32Bit(request.OldPassword)}
	if err := user.LoginUser(); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR) + ":" + msg.GetMsg(tcode.USERNAME_OR_PASSWORD_ERROR)
		return
	}
	args := make(map[string]interface{})
	args["password"] = MD5Helper.EncryptMD5To32Bit(request.NewPassword)
	if err := user.UpdataUser(args); err != nil {
		responseData.Message = msg.GetMsg(tcode.UPDATE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.UPDATE_SUCCESS)
	return
}

// 添加用户地址信息
func AddUserInfo(token string, request entity.AddUserInfo) (responseData entity.ResponseData) {
	user := models.User{Token: token}
	if !user.QueryUser() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + "，用户不存在"
		return
	}
	userInfo := models.UserInfo{
		Provice:     request.Provice,
		City:        request.City,
		ShopAddress: request.ShopAddress,
		UserID:      user.ID,
	}
	if err := userInfo.AddUserInfo(); err != nil {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.ADD_SUCCESS)
	return
}

// 查询用户所有地址信息
func QueryUserInfoByUserID(tel string) (responseData entity.ResponseData) {
	user := models.User{}
	user.Tel = tel
	if !user.QueryUser() {
		responseData.Message = msg.GetMsg(tcode.ADD_ERROR) + ",该用户不存在"
		return
	}
	userInfo := models.UserInfo{}
	userInfo.UserID = user.ID
	userInfos := userInfo.QueryUserInfoByUserID()
	if len(userInfos) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	data := make(map[string]interface{})
	data["userInfo"] = userInfos
	responseData.Data = data
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	return
}

// 查询用户地址信息详情
func QueryUserInfoByID(id int64) (responseData entity.ResponseData) {
	userInfo := models.UserInfo{}
	userInfo.ID = id
	if err := userInfo.QueryUserInfoByID; err != nil {
		responseData.Message = msg.GetMsg(tcode.QUERY_ERROR)
		return
	}
	data := make(map[string]interface{})
	data["userInfo"] = userInfo
	responseData.Data = data
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	return
}

// 删除用户地址信息
func DeleteUserInfo(ids []int64) (responseData entity.ResponseData) {
	if len(ids) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTNULL)
		return
	}
	userInfo := models.UserInfo{}
	if err := userInfo.DeleteUserInfo(ids); err != nil {
		responseData.Message = msg.GetMsg(tcode.DELETE_ERROR)
		return
	}
	responseData.Status = true
	responseData.Message = msg.GetMsg(tcode.DELETE_SUCCESS)
	return
}

// 查询用户（分页） admin
func QueryUserByLimitOffset(pageSize int, page int) (responseData entity.ResponseData) {
	users := models.QueryUserByLimitOffset(pageSize, page)
	responseData.Message = msg.GetMsg(tcode.QUERY_SUCCESS)
	if len(users) == 0 {
		responseData.Message = msg.GetMsg(tcode.NOTMORE)
	}
	count := models.QueryUserCount()
	data := make(map[string]interface{})
	data["users"] = users
	data["count"] = count
	responseData.Data = data
	responseData.Status = true
	return
}
