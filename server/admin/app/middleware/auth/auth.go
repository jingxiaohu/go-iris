package auth

import (
	"goiris/common"
	"goiris/common/icasbin"
)

func InitAuth() {
	icasbin.InitCasbin(common.G_AppConfig.Domain)
	icasbin.G_Casbin.AddFunction("myFunc", myFunc)
}
func myFunc(args ...interface{}) (interface{}, error) {
	match := func(key string) bool {
		if key == "/api/admin/user/logout" || key == "/api/admin/user/userBaseInfo" ||
			key == "/api/admin/menu/all" || key == "/api/admin/menu/own" {
			return true
		}
		return false
	}(args[0].(string))
	return (bool)(match), nil
}
