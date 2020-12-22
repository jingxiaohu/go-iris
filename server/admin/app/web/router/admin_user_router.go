package router

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"goiris/admin/app/middleware/jwt"
	"goiris/admin/app/web/mapper"
	"goiris/admin/app/web/vo"
	"goiris/common"
	"goiris/common/icasbin"
	"goiris/common/model"
	"goiris/common/storage"
	"goiris/common/support"
	"strconv"
	"strings"
	"time"
)

func adminUserRouter(party iris.Party) {
	var (
		adinUser = party.Party("/user")
		aus      = AdminUserService{mapper: mapper.AdminUserMapper{}}
	)
	adinUser.Post("/", hero.Handler(aus.create))
	adinUser.Delete("/", hero.Handler(aus.delete))
	adinUser.Put("/", hero.Handler(aus.put))
	adinUser.Post("/table", hero.Handler(aus.table))
	//
	adinUser.Post("/login", hero.Handler(aus.login))
	adinUser.Get("/userBaseInfo", hero.Handler(aus.userBaseInfo))
	adinUser.Get("/logout", hero.Handler(aus.logout))
}

type AdminUserService struct {
	mapper mapper.AdminUserMapper
}

func (aus *AdminUserService) create(ctx iris.Context) {
	var (
		code      support.Code
		err       error
		vo = new(vo.AcceptAdminUserVO)
	)
	if err = ctx.ReadJSON(&vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}

	if vo.AdminUser.Gender == 0 {
		vo.AdminUser.Gender = 2 // 默认性别设置为女
	}
	vo.AdminUser.Password = support.AESEncrypt([]byte(vo.AdminUser.Password))
	vo.AdminUser.Enable = true
	vo.AdminUser.CreateTime = time.Now()
	if err = aus.mapper.Insert(vo); err != nil {
		code = support.CODE_USER_PHONE_DUPLICATE
		goto ERR
	}
	icasbin.G_Casbin.LoadPolicy()
	support.Ok_(ctx, support.CODE_USER_REGISTE_OK)
	return
ERR:
	golog.Errorf("用户[%s]注册失败。原因：%s，错误：%s", vo.AdminUser.Phone, code, err)
	support.InternalServerError(ctx, code)
}

func (aus *AdminUserService) delete(ctx iris.Context) {
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
	if err = aus.mapper.Delete(ids); err != nil {
		goto ERR
	}
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("操作失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

func (aus *AdminUserService) put(ctx iris.Context) {
	var (
		code      support.Code
		err       error
		vo = new(vo.AcceptAdminUserVO)
	)
	if err = ctx.ReadJSON(&vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	vo.AdminUser.Password = support.AESEncrypt([]byte(vo.AdminUser.Password))
	vo.AdminUser.UpdateTime = time.Now()
	if err = aus.mapper.UpdateOne(vo); err != nil {
		goto ERR
	}
	support.Ok_(ctx, support.CODE_OK)
	return
ERR:
	golog.Errorf("更新用户[%d]信息失败。原因：%s，错误：%s", vo.AdminUser.Id, code.String(), err)
	support.InternalServerError(ctx, code)
}

func (aus *AdminUserService) table(ctx iris.Context) {
	var (
		err    error
		code   support.Code
		vo     = new(vo.AdminUserVO)
		total  int64
		result []*model.AdminUser
	)
	if err = ctx.ReadJSON(vo); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	vo.Init()
	if total, result, err = aus.mapper.FindList(vo); err != nil {
		goto ERR
	}
	support.PaginationTableData(ctx, total, result)
	return
ERR:
	golog.Errorf("失败。原因：%s，错误：%s", code.String(), err)
	support.InternalServerError(ctx, code)
}

// 登录获取token
func (aus *AdminUserService) login(ctx iris.Context) {
	var (
		code       support.Code
		err        error
		adminUser  = new(model.AdminUser)
		mAdminUser = new(model.AdminUser)
		ckPassword bool
		token      string
	)
	// TODO 同时登陆一个账号，将已登陆的用户挤下线
	if err = ctx.ReadJSON(&adminUser); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}
	if err = support.G_validate.Struct(adminUser); err != nil {
		code = support.CODE_SYS_PARSE_PARAMS_ERROR
		goto ERR
	}

	mAdminUser.Username = adminUser.Username
	mAdminUser.Enable = true
	if mAdminUser, err = aus.mapper.FindOne(mAdminUser); err != nil {
		code = support.CODE_USER_PHONE_FAILUR
		goto ERR
	}
	ckPassword = support.CheckPWD(adminUser.Password, mAdminUser.Password)
	if !ckPassword {
		code = support.CODE_USER_PASSWORD_FAILUR
		goto ERR
	}

	if token, err = jwt.G_JWT.GenerateToken(mAdminUser); err != nil {
		code = support.CODE_TOKEN_CREATE_FAILUR
		goto ERR
	}
	if err = storage.G_Redis.SetToken(support.REDIS_ADMIN_FORMAT, strconv.FormatInt(mAdminUser.Id, 10), token); err != nil {
		code = support.CODE_TOKEN_CACHE_ERROR
		goto ERR
	}
	support.Ok(ctx, support.CODE_USER_LOGIN_OK, iris.Map{
		"token":  token,
		"expire": common.G_AppConfig.Own.JWTTimeout,
	})
	return
ERR:
	golog.Errorf("~~> admin后台用户[%s]登录失败。原因：%s，错误：%s", adminUser.Username, code.String(), err)
	support.InternalServerError(ctx, code)
}

// 登录后获取用户信息
func (aus *AdminUserService) userBaseInfo(ctx iris.Context) {
	var (
		code      support.Code
		err       error
		adminUser *model.AdminUser
		uid       int64
	)
	if uid, err = ctx.Values().GetInt64(support.UID); err != nil {
		goto ERR
	}
	golog.Infof("userBaseInfo, uid=%d", uid)
	if adminUser, err = aus.mapper.FindOne(&model.AdminUser{Id: uid,}); err != nil {
		code = support.CODE_SYS_FAILUR
		goto ERR
	}
	adminUser.Password = ""
	adminUser.Online = true
	support.Ok(ctx, support.CODE_USER_LOGIN_OK, adminUser)
	return
ERR:
	golog.Errorf("~~> 获取用户[%d]基础信息失败。原因：%s， 错误：%s", uid, code.String(), err.Error())
	support.InternalServerError(ctx, code)
	return
}

func (aus *AdminUserService) logout(ctx iris.Context) {
	var (
		err  error
		code support.Code
		uid  int64
	)
	if uid, err = ctx.Values().GetInt64(support.UID); err != nil {
		goto ERR
	}
	if _, err = storage.G_Redis.DelToken(support.REDIS_ADMIN_FORMAT, strconv.FormatInt(uid, 10)); err != nil {
		code = support.CODE_TOKEN_CACHE_ERROR
		goto ERR
	}
	support.Ok_(ctx, support.Code(200))
	return
ERR:
	golog.Errorf("~~> 用户[%d]退出失败。原因：%s，错误：%s", uid, code.String(), err)
	support.InternalServerError(ctx, code)
	return
}
