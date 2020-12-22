package storage

import (
	"encoding/json"
	"goiris/admin/app/bindata/conf"
	"goiris/admin/app/middleware/auth"
	"goiris/common"
	"goiris/common/icasbin"
	"goiris/common/model"
	"testing"
	"time"
)

func Test_menu(t *testing.T) {
	common.InitConfig(conf.Asset)
	InitGorm()
	var (
		err    error
		result = make([]*model.Menu, 0)
	)
	if err = G_DB.Preload("Children.Meta").Preload("Meta").Find(&result, "is_sub = 0 AND enable = 1").Error; err != nil {
		t.Fatal(err)
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", result)
	t.Logf("%s", bytes)
	t.Logf("len=%d\n", len(result))
}

func Test_a1(t *testing.T) {
	common.InitConfig(conf.Asset)
	auth.InitAuth()
	is := icasbin.G_Casbin.Enforce("90", "a", "/api/admin/role/table", "POST", ".*")
	t.Logf("%v", is)
	//t.Log(icasbin.G_Casbin.GetPolicy())
	//t.Log(icasbin.G_Casbin.GetAllSubjects())
	//t.Log(icasbin.G_Casbin.GetAllObjects())
	//t.Log(icasbin.G_Casbin.GetAllActions())
	//t.Log(icasbin.G_Casbin.GetAllRoles())
	//t.Log(icasbin.G_Casbin.GetPermissionsForUser("90"))
	//t.Log(icasbin.G_Casbin.GetImplicitRolesForUser("90", common.G_AppConfig.Domain))
	icasbin.G_Casbin.AddRoleForUserInDomain("5", "admin", "a")
}

func Test_menu1(t *testing.T) {
	common.InitConfig(conf.Asset)
	//InitGorm()
	auth.InitAuth()
	//icasbin.InitCasbin("a")
	var (
		err error
		res []string
	)
	res, err = icasbin.G_Casbin.GetRolesForUser("90")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", res)

	res = icasbin.G_Casbin.GetImplicitRolesForUser("90")
	t.Logf("%v", res)

	isPass := icasbin.G_Casbin.Enforce("90a", "a", "/admin/bb", "POST", ".*")
	t.Log(isPass)

	t.Log(icasbin.G_Casbin.GetModel())
	t.Log(icasbin.G_Casbin.GetPolicy())
	t.Log(icasbin.G_Casbin.GetGroupingPolicy())
	t.Log(icasbin.G_Casbin.GetAllSubjects())
	t.Log(icasbin.G_Casbin.GetAllObjects())
	t.Log(icasbin.G_Casbin.GetAllActions())
	t.Log(icasbin.G_Casbin.GetAllRoles())
}

func Test_menu2(t *testing.T) {
	common.InitConfig(conf.Asset)
	InitGorm()
	var (
		err   error
		resut = make([]*model.CasbinRule, 0)
	)
	if err = G_DB.Preload("Mids").First(&resut, "p_type='g'").Error; err != nil {
		t.Log(err)
	}
	t.Logf("%v", resut)
}

func Test_menu3(t *testing.T) {
	common.InitConfig(conf.Asset)
	auth.InitAuth()
	t.Log(time.Now().String())
	a := icasbin.G_Casbin.AddPolicy("demo", "a", "/api.", "GET", ".*", "测试", time.Now().String())
	t.Log(a)
	icasbin.G_Casbin.DeleteRole("demo")
}
