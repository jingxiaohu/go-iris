package model

import (
	"time"
)

// 部门表
type Dep struct {
	Id         int64        `gorm:"primary_key" json:"id" form:"id"`
	DepName    string       `xorm:"varchar(64) notnull" json:"depName" form:"depName"`
	Leader     string       `xorm:"varchar(64) notnull" json:"leader" form:""leader`
	Phone      string       `xorm:"varchar(64) notnull" json:"phone" form:"phone"`
	Email      string       `xorm:"varchar(64) notnull" json:"email" form:"email"`
	Disabled   bool         `xorm:"notnull tinyint(0)" json:"disabled" form:"disabled"` //0:可用 否则:不可用
	ParentId   int64        `xorm:"INT(10) notnull" json:"parentId" form:"parentId"`
	DepDesc    string       `xorm:"varchar(255) notnull" json:"depDesc" form:"depDesc"`
	CreateTime time.Time    `json:"createTime"`
	UpdateTime time.Time    `json:"updateTime" form:"updateTime"`
	SubDeps    []*Dep       `json:"subDeps" gorm:"foreignkey:ParentId;"`
	Admins     []*AdminUser `json:"admins" gorm:"many2many:dep_admin;"`
}
