package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"goiris/admin/app/web/mapper"
	"goiris/admin/app/web/vo"
	"goiris/common/icasbin"
	"goiris/common/model"
	"goiris/common/support"
	"strconv"
	"strings"
)

func policyRouter(party iris.Party) {
	policy := party.Party("/policy")
	ps := PolicyService{mapper: &mapper.PolicyMapper{}}

	policy.Post("/", hero.Handler(ps.create))
	policy.Delete("/", hero.Handler(ps.delete))
	policy.Put("/", hero.Handler(ps.put))
	policy.Post("/table", hero.Handler(ps.table)) // 权限管理报表
	policy.Get("/all", hero.Handler(ps.all))   // 所有权限
}

type PolicyService struct {
	mapper *mapper.PolicyMapper
}

func (ps *PolicyService) create(ctx iris.Context) {
	var (
		code support.Code
		err  error
		role = new(model.CasbinRule)
	)
	if err = ctx.ReadJSON(&role); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(role); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if !icasbin.G_Casbin.AddPolicy(role.V0, role.V1, role.V2, role.V3, role.V4, role.V5) {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	support.Ok_(ctx, support.CODE_USER_REGISTE_OK)
	return
ERR:
	golog.Errorf("创建角色[%s]失败。原因：%s，错误：%s", role.V5, code, err)
	support.InternalServerError(ctx, code)
}

func (ps *PolicyService) delete(ctx iris.Context) {
	var (
		code support.Code
		err  error
		ids  []int64
	)
	if ids, err = func() (result []int64, err error) {
		idStr := strings.Split(ctx.URLParam("ids"), ",")
		for _, v := range idStr {
			var id int64
			if id, err = strconv.ParseInt(v, 10, 64); err != nil {
				code = support.CODE_SYS_PARSE_PARAMS_ERROR
				return nil, err
			}
			result = append(result, id)
		}
		return
	}(); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = ps.mapper.Delete(ids); err != nil {
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("操作失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (ps *PolicyService) put(ctx iris.Context) {
	var (
		code support.Code
		err  error
		role = new(model.CasbinRule)
	)
	if err = ctx.ReadJSON(&role); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(role); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = ps.mapper.UpdateOne(role); err != nil {
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("更新角色[%s]信息失败。原因：%s，错误：%s", role.V5, code.String(), err)
	support.InternalServerError(ctx, code)
}

func (ps *PolicyService) table(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		vo     = new(vo.RoleVO)
		total int64
		result []*model.CasbinRule
	)
	if err = ctx.ReadJSON(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	vo.Init()
	if total, result, err = ps.mapper.Table(vo); err != nil {
		goto ERR
	}
	support.PaginationTableData(ctx, total, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (ps *PolicyService) all(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		result []*model.CasbinRule
	)
	if result, err = ps.mapper.All(); err != nil {
		goto ERR
	}
	support.Ok(ctx, support.CODE_OK, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}
