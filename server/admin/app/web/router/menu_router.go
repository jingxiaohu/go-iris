package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"goiris/admin/app/web/mapper"
	"goiris/admin/app/web/vo"
	"goiris/common/model"
	"goiris/common/support"
	"strconv"
	"strings"
	"time"
)

func menuRouter(party iris.Party) {
	menu := party.Party("/menu")
	mm := MenuService{mapper: mapper.MenuMapper{}}

	menu.Get("/all", hero.Handler(mm.all)) // 获取所有菜单
	menu.Get("/own", hero.Handler(mm.own)) // 获取自己应有的菜单id

	menu.Post("/", hero.Handler(mm.create))
	menu.Delete("/", hero.Handler(mm.delete))
	menu.Put("/", hero.Handler(mm.put))
	menu.Post("/table", hero.Handler(mm.table))
	menu.Get("/parentMenus", hero.Handler(mm.parentMenus)) // 获取所有的父菜单，供菜单修改时选择
}

type MenuService struct {
	mapper mapper.MenuMapper
}

// 动态菜单实现：账号登录时，大权限账号先登录后退出，小权限账号再登陆的情况下，由于vue不能清空路由，会导致前端看见不属于自己的菜单且会提示duplicate重复路由。解决如下：
// 1. 前端路由控制：将所有路由菜单数据发给前端，前端iview-admin通过meta中HideInMenu属性控制菜单显示
// 2. 再加上后台casbin控制
func (ms *MenuService) all(ctx iris.Context) {
	var (
		code  support.Code
		err   error
		uid   = ctx.Values().GetString(support.UID)
		menus []*model.Menu
		owned []int64
		have  func(id int64) bool
	)
	if menus, err = ms.mapper.AllMenu(); err != nil {
		code = support.CODE_FAILUR
		goto ERR
	}
	// 属于自己的菜单的id
	if owned, err = ms.mapper.OwnMenu(uid); err != nil {
		code = support.CODE_FAILUR
		goto ERR
	}
	// 是否在关联的角色菜单中，true=有该菜单
	have = func(id int64) bool {
		for _, vv := range owned {
			if id == vv {
				return true
			}
		}
		return false
	}
	// 将不属于自己的菜单设置隐藏
	for _, v1 := range menus {
		if v1.Meta == nil {
			v1.Meta = &model.MenuMeta{}
		}
		if !v1.Meta.HideInMenu {
			if have(v1.Id) {
				v1.Meta.HideInMenu = false // 显示菜单
			} else {
				v1.Meta.HideInMenu = true
			}
		}
		// 判断子菜单
		for _, v2 := range v1.Children {
			if v2.Meta == nil {
				v2.Meta = &model.MenuMeta{}
			}
			if !v2.Meta.HideInMenu {
				if have(v2.Id) {
					v2.Meta.HideInMenu = false // 有子菜单
					v1.Meta.HideInMenu = false // 同时父菜单也设置显示
				} else {
					v2.Meta.HideInMenu = true
				}
			}
		}
		/*
			// true=该菜单的所有子菜单都隐藏着
			var allChildHide = func() bool {
				for _, v3 := range v1.Children {
					if v3.Meta.HideInMenu == false {
						return false
					}
				}
				return true
			}
			// 如果父菜单显示，而对应的子菜单却没有，则该父菜单不显示
			if v1.Meta.HideInMenu == false && (v1.Children == nil || allChildHide()) {
				v1.Children = make([]*model.Menu, 0)
				v1.Meta.HideInMenu = true
			}
		*/
	}
	support.Ok(ctx, support.CODE_OK, menus)
	return
ERR:
	golog.Errorf("~~> 用户[%s]获取所有菜单失败。原因：%s， 错误：%s", uid, code.String(), err)
	support.InternalServerError(ctx, code)
	return
}

func (ms *MenuService) own(ctx iris.Context) {
	var (
		code  support.Code
		err   error
		uid   = ctx.Values().GetString(support.UID)
		owned []int64
	)
	if owned, err = ms.mapper.OwnMenu(uid); err != nil {
		code = support.CODE_FAILUR
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, owned)
	return
ERR:
	golog.Errorf("~~> 用户[%s]获取自己的菜单id失败。原因：%s， 错误：%s", uid, code.String(), err)
	support.InternalServerError(ctx, code)
	return
}

func (ms *MenuService) create(ctx iris.Context) {
	var (
		code support.Code
		err  error
		menu = new(model.Menu)
	)
	if err = ctx.ReadJSON(&menu); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(menu); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = ms.mapper.Insert(menu); err != nil {
		goto ERR
	}
	support.Ok_(ctx, support.CODE_USER_REGISTE_OK)
	return
ERR:
	golog.Errorf("新建菜单[%s]失败。原因：%s，错误：%s", menu.Name, code, err)
	support.InternalServerError(ctx, code)
}

func (ms *MenuService) delete(ctx iris.Context) {
	var (
		code  support.Code
		err   error
		mids  []int64 // 菜单id
		mMids []int64 // 菜单对应的元信息的id
	)
	if mids, mMids, err = func() (result1, result2 []int64, err error) {
		midStr := strings.Split(ctx.URLParam("mids"), ",")
		mMidStr := strings.Split(ctx.URLParam("mMids"), ",")
		for _, v := range midStr {
			var id int64
			if id, err = strconv.ParseInt(v, 10, 64); err != nil {
				code = support.CODE_SYS_PARSE_PARAMS_ERROR
				return nil, nil, err
			}
			result1 = append(result1, id)
		}
		for _, v := range mMidStr {
			var id int64
			if id, err = strconv.ParseInt(v, 10, 64); err != nil {
				code = support.CODE_SYS_PARSE_PARAMS_ERROR
				return nil, nil, err
			}
			result2 = append(result2, id)
		}
		return
	}(); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = ms.mapper.Delete(mids, mMids); err != nil {
		goto ERR
	}
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("操作失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (ms *MenuService) put(ctx iris.Context) {
	var (
		code support.Code
		err  error
		menu = new(model.Menu)
	)
	if err = ctx.ReadJSON(&menu); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(menu); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	menu.UpdateTime = time.Now()
	if err = ms.mapper.UpdateOne(menu); err != nil {
		goto ERR
	}
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("更新菜单[%d]信息失败。原因：%s，错误：%s", menu.Id, code.String(), err)
	support.InternalServerError(ctx, code)
}

func (ms *MenuService) table(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		vo     = new(vo.MenuVO)
		total  int64
		result []*model.Menu
	)
	if err = ctx.ReadJSON(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	vo.Init()
	if total, result, err = ms.mapper.FindList(vo); err != nil {
		goto ERR
	}
	support.PaginationTableData(ctx, total, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (ms *MenuService) parentMenus(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		result []*model.Menu
	)
	if result, err = ms.mapper.ParentMenus(); err != nil {
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}
