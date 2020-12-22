package model

// 部门-用户关联表
type DepUser struct {
	Id  int64 `xorm:"pk autoincr INT(10) notnull" json:"id"`
	Rid int64 `xorm:"pk autoincr INT(10) notnull" json:"rid"`
	Mid int64 `xorm:"pk autoincr INT(10) notnull" json:"mid"`
}
