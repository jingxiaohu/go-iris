package vo

import (
	"goiris/common/model"
	"goiris/common/support"
)

type RoleVO struct {
	support.Pagination
	Domain string `json:"domain"`
}

// 更新和创建角色时需要的vo
type AcceptRoleVO struct {
	Role       *model.CasbinRule   `json:"role"`       // 角色本身的信息
	AddMids    []int64             `json:"addMids"`    // 增加的角色菜单
	DeleteMids []int64             `json:"deleteMids"` // 需删除的角色菜单
}