package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"goiris/admin/app/web/mapper"
	"goiris/admin/app/web/vo"
	"goiris/common"
	"goiris/common/icasbin"
	"goiris/common/model"
	"goiris/common/support"
	"strings"
)

func roleRouter(party iris.Party) {
	role := party.Party("/role")
	mm := RoleService{mapper: &mapper.RoleMapper{}}

	role.Post("/", hero.Handler(mm.create))
	role.Delete("/", hero.Handler(mm.delete))
	role.Put("/", hero.Handler(mm.put))
	role.Get("/domainList", hero.Handler(mm.domainList))
	role.Get("/group", hero.Handler(mm.group))                                // 获取所有:用户组角色列表
	role.Post("/table", hero.Handler(mm.table))                               // 后台系统的权限报表
	role.Get("/roleOfMenus/{roleKey:string}", hero.Handler(mm.roleOfMenus))   // 角色拥有的对应菜单集合
	role.Get("/roleOfPolicys/{roleKey:string}/{domain:string}", hero.Handler(mm.roleOfPolicys)) // 角色拥有的对应权限集合
}

type RoleService struct {
	mapper *mapper.RoleMapper
}

func (rs *RoleService) create(ctx iris.Context) {
	var (
		code support.Code
		err  error
		vo   = new(vo.AcceptRoleVO)
	)
	if err = ctx.ReadJSON(&vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if vo.Role.PType != "g" {
		vo.Role.PType = "g"
	}
	if err = rs.mapper.Insert(vo); err != nil {
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	//golog.Info(icasbin.G_Casbin.LoadPolicy().Error())
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("创建角色[%s]失败。原因：%s，错误：%s", vo.Role.V5, code, err)
	support.InternalServerError(ctx, code)
}

func (rs *RoleService) delete(ctx iris.Context) {
	var (
		code     support.Code
		err      error
		roleKeys = strings.Split(ctx.URLParam("roleKeys"), ",")
	)
	if err = rs.mapper.Delete(roleKeys); err != nil {
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("操作失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (rs *RoleService) put(ctx iris.Context) {
	var (
		code support.Code
		err  error
		vo   = new(vo.AcceptRoleVO)
	)
	if err = ctx.ReadJSON(&vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = rs.mapper.UpdateOne(vo); err != nil {
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("更新角色[%s]信息失败。原因：%s，错误：%s", vo.Role.V5, code.String(), err)
	support.InternalServerError(ctx, code)
}

func (rs *RoleService) domainList(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		result []string
	)
	if result, err = rs.mapper.DomainList(); err != nil {
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (rs *RoleService) group(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		result []*model.CasbinRule
	)
	if result, err = rs.mapper.Group(); err != nil {
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (rs *RoleService) table(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		vo     = new(vo.RoleVO)
		result []*model.CasbinRule
	)
	if err = ctx.ReadJSON(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	vo.Init()
	if result, err = rs.mapper.Table(vo); err != nil {
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (rs *RoleService) roleOfMenus(ctx iris.Context, roleKey string) {
	var (
		code   support.Code
		err    error
		result []*model.RoleMenu
	)
	if result, err = rs.mapper.RoleOfMenus(roleKey); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, result)
	return
ERR:
	golog.Errorf("操作失败。原因：%s，错误：%s", code, err)
	support.InternalServerError(ctx, code)
}

func (rs *RoleService) roleOfPolicys(ctx iris.Context, roleKey, domain string) {
	var (
		code   support.Code
		err    error
		result []*model.CasbinRule
	)
	if domain == "" {
		domain = common.G_AppConfig.Domain
	}
	if result, err = rs.mapper.RoleOfPolicys(roleKey, domain); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, result)
	return
ERR:
	golog.Errorf("操作失败。原因：%s，错误：%s", code, err)
	support.InternalServerError(ctx, code)
}
