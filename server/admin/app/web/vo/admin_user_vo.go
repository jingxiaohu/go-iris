package vo

import (
	"goiris/common/model"
	"goiris/common/support"
)

type AdminUserVO struct {
	support.Pagination
}

// 更新和创建用户时需要的vo
type AcceptAdminUserVO struct {
	AdminUser  *model.AdminUser    `json:"adminUser"`  // 用户本身的信息
	AddRoles   []*model.CasbinRule `json:"addRoles"`   // 增加的角色
	DeleteRids []int64             `json:"deleteRids"` // 需删除的角色
}
