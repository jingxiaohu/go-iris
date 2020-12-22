package model

// 角色和权限表
type CasbinRule struct {
	Id         int64         `gorm:"primary_key" json:"id"`
	PType      string        `json:"pType"`
	V0         string        `json:"v0"`
	V1         string        `json:"v1"`
	V2         string        `json:"v2"`
	V3         string        `json:"v3"`
	V4         string        `json:"v4"`
	V5         string        `json:"v5"` // 名称
	MenuList   []*RoleMenu   `json:"menuList" gorm:"-"`
	PolicyList []*CasbinRule `json:"policyList" gorm:"-"`
	//V6    string `json:"v6"` // 菜单id。后台系统，菜单和权限一对一关系
}

func (c *CasbinRule) TableName() string {
	return "casbin_rule" //as Gorm keeps table names are plural, and we love consistency
}
