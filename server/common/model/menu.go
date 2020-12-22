package model

import (
	"time"
)

type (
	// 菜单表
	Menu struct {
		Id         int64     `json:"id" gorm:"primary_key"`
		Path       string    `json:"path"`
		Modular    string    `json:"modular"`
		Component  string    `json:"component"`
		Name       string    `json:"name"`
		ParentId   int64     `json:"parentId"`
		IsSub      bool      `json:"isSub"`
		Sort       int       `json:"sort"`
		CreateTime time.Time `json:"createTime"`
		UpdateTime time.Time `json:"updateTime"`
		Children   []*Menu   `json:"children" gorm:"FOREIGNKEY:ParentId;"`
		MetaId     int64     `json:"metaId"`
		Meta       *MenuMeta `json:"meta" gorm:"FOREIGNKEY:MetaId;"`
	}
	// 路由可配项
	// 没有绑定每个用户的配置，则对系统所有用户生效
	MenuMeta struct {
		Id int64 `json:"id" gorm:"primary_key"`
		// { String|Number|Function } 菜单显示的名称
		// 显示在侧边栏、面包屑和标签栏的文字
		// 使用'{{ 多语言字段 }}'形式结合多语言使用，例子看多语言的路由配置;
		// 可以传入一个回调函数，参数是当前路由对象，例子看动态路由和带参路由
		Title           string   `json:"title" `
		HideInBread     bool     `json:"hideInBread"`     // (default: false) 设为true后此级路由将不会出现在面包屑中，示例看QQ群路由配置
		HideInMenu      bool     `json:"hideInMenu"`      // (default: false) 设为true后在左侧菜单不会显示该页面选项
		ShowAlways      bool     `json:"showAlways"`      // (default: false) 设为true后如果该路由只有一个子路由，在菜单中也会显示该父级菜单
		NotCache        bool     `json:"notCache"`        // (default: false) 设为true后页面在切换标签后不会缓存，如果需要缓存，无需设置这个字段，而且需要设置页面组件name属性和路由配置的name一致
		Access          string `json:"access"`          // (default: null) 可访问该页面的权限数组，当前路由设置的权限会影响子路由
		Icon            string   `json:"icon"`            // (default: null) 该页面在左侧菜单、面包屑和标签导航处显示的图标，如果是自定义图标，需要在图标名称前加下划线'_'
		BeforeCloseName string   `json:"beforeCloseName"` // (default: null) 设置该字段，则在关闭当前tab页时会去'@/router/before-close.js'里寻找该字段名对应的方法，作为关闭前的钩子函数
		Href            string   `json:"href"`            // (default: null) 用于跳转到外部连接
	}
)
