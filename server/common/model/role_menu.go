package model

// 角色-菜单关联表
type RoleMenu struct {
	//Id  int64 `xorm:"pk autoincr INT(10) notnull" json:"id"`
	RoleKey string `gorm:"primary_key" json:"roleKey"`
	Mid     int64  `gorm:"primary_key" json:"mid"`
}
