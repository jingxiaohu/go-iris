package icasbin

import (
	"github.com/casbin/casbin"
	"github.com/kataras/golog"

	"goiris/common"
)

var G_Casbin *casbin.Enforcer

func InitCasbin(domain string) {
	var (
		err    error
		config = common.G_DBConfig
		url    = config.DBConnUrl()
	)
	modelStr := `
		[request_definition]
		r = sub, dom, obj, act, suf
		
		[policy_definition]
		p = sub, dom, obj, act, suf, name
		
		[role_definition]
		g = _, _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act) && regexMatch(r.suf, p.suf)`
	if domain == "a" {
		modelStr += ` || myFunc(r.obj) || r.sub == "1"` // 设置id=1为超级用户
	}
	model := casbin.NewModel(modelStr)
	e := casbin.NewEnforcer(model, NewAdapter(config.Mysql.Dialect, url, true))
	e.EnableLog(true)
	if err = e.LoadFilteredPolicy(Filter{
		PType: []string{"p", "g"},
		V1:    []string{domain},
		V2:    []string{domain},
	}); err != nil {
		golog.Fatalf("~~> Casbin策略加载失败,原因:%s", err)
	}
	G_Casbin = e
}
