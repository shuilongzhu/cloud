package init_data

import (
	"cloud/common/db"
	"cloud/common/pojo"
	"cloud/common/utils/data_convert"
	"cloud/error_code"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"strings"
	"time"
)

// CommonInit @description: 项目默认数据初始化
func CommonInit() {
	go addDefaultInfo()
	go db.DataInit()
	pojo.ClientIp = GetClientIp()
}

// addDefaultInfo @description: 确保项目有一个admin超级管理员用户，未分组，及其设备未分组存在
// @return code
func addDefaultInfo() (code int) {
	//查询超级管理员初始用户组
	code, userGroup := db.UserGroupSelect(pojo.AdminDefaultUserGroupId)
	if error_code.Success != code {
		return
	}

	//确保admin用户初始用户组存在
	if code = addAdminUserGroup(userGroup); error_code.Success != code {
		return
	}

	//查询admin用户
	code, userInfo := db.UserInfoSelect(pojo.AdminUserName)
	if error_code.Success != code {
		return
	}

	//确保admin用户存在
	if code = addAdminUserInfo(userInfo); error_code.Success != code {
		return
	}

	//查询admin用户设备未分组信息
	code, deviceGroup := db.DeviceGroupSelect(pojo.AdminDefaultDeviceGroupId)
	if error_code.Success != code {
		return
	}

	//确保admin用户有一个默认设备未分组
	code = addAdminDeviceGroup(deviceGroup)
	return
}

// addAdminUserGroup @description: 确保admin用户初始用户组存在
// @parameter userGroup
// @return code
func addAdminUserGroup(userGroup pojo.UserGroup) (code int) {
	code = error_code.Success

	if 0 == userGroup.Id { //不存在
		log.Info("Add default admin group")
		userGroup.UserGroupId = pojo.AdminDefaultUserGroupId
		userGroup.UserGroupName = "系统初始用户组"
		userGroup.UserGroupNote = "系统初始用户组，后面可以修改名称（建议公司名称）"
		userGroup.CreatedTime = time.Now()
		userGroup.UpdateTime = time.Now()
		//创建超级管理员初始用户组
		if code = db.UserGroupCreate(userGroup); error_code.Success != code {
			log.Errorln("新增系统初始用户admin初始用户组失败！")
			return
		}

		log.Println("新增系统初始用户admin初始用户组成功！UserGroupId:", userGroup.UserGroupId)
	}

	log.Println("admin用户初始用户组已存在，UserGroupId:", pojo.AdminDefaultUserGroupId)
	return
}

// addAdminUserInfo @description: 确保admin用户存在
// @parameter userInfo
// @return code
func addAdminUserInfo(userInfo pojo.UserInfo) (code int) {
	code = error_code.Success

	if 0 == userInfo.Id { //不存在，添加系统初始用户admin
		userInfo.UserName = pojo.AdminUserName
		userInfo.UserId = pojo.IntPrefix + strconv.FormatInt(data_convert.NextId(), 10)
		userInfo.UserGroupId = pojo.AdminDefaultUserGroupId
		userInfo.Nickname = pojo.AdminUserName
		userInfo.UserPassword = data_convert.GetMD5Encode(pojo.AdminPassword)
		userInfo.Role = 1 //角色类型 0：普通用户；1：系统超级管理员；2:运维观察员；3:模型管理员
		userInfo.UserGroupName = "系统初始用户组"
		userInfo.Note = "系统初始用户admin"
		userInfo.CreatedTime = time.Now()
		userInfo.UpdateTime = time.Now()
		userInfo.LoginTime = time.Now()
		userInfo.UserState = 1 //用户状态 1：正常；2：已禁用；
		//创建admin用户
		if code = db.UserInfoCreate(userInfo); error_code.Success != code {
			log.Errorln("新增系统初始用户admin失败！")
			return
		}
		log.Println("新增系统初始用户admin成功！UserId:", userInfo.UserId)
	}
	log.Println("admin用户已存在，UserId:", userInfo.UserId)

	pojo.AdminUserId = userInfo.UserId
	return
}

// addAdminDeviceGroup @description: 确保admin用户有一个默认未分组
// @parameter deviceGroup
// @return code
func addAdminDeviceGroup(deviceGroup pojo.DeviceGroup) (code int) {
	code = error_code.Success

	if 0 == deviceGroup.Id { //不存在，给admin用户添加一个默认未分组
		deviceGroup.Id = data_convert.NextId()
		deviceGroup.DeviceGroupId = pojo.AdminDefaultDeviceGroupId
		deviceGroup.DeviceGroupName = "未分组设备"
		deviceGroup.UserGroupId = pojo.AdminDefaultUserGroupId
		deviceGroup.DeviceGroupNote = "未分组设备"
		deviceGroup.CreatedTime = time.Now()
		deviceGroup.UpdateTime = time.Now()
		//创建admin用户设备未分组
		if code = db.DeviceGroupCreate(deviceGroup); error_code.Success != code {
			log.Errorln("新增admin用户设备未分组失败！")
			return
		}
		log.Println("新增admin用户设备未分组成功！DeviceGroupId:", pojo.AdminDefaultDeviceGroupId)
	}
	log.Println("admin用户设备未分组已存在，DeviceGroupId:", pojo.AdminDefaultDeviceGroupId)
	return
}

// GetClientIp @description: 获取服务IP地址
// @return string
func GetClientIp() string {
	var str string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return str
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						var tempStr = ipnet.IP.String()
						var strList = strings.Split(tempStr, ".")
						if "1" != strList[len(strList)-1] { //服务器ip地址最后一位不为1
							str = tempStr
						}
					}
				}
			}
		}
	}
	log.Println("service ip is ", str)
	return str
}
