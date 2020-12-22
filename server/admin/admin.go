package main

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"goiris/admin/app/bindata/conf"
	"goiris/admin/app/bindata/static"
	"goiris/admin/app/middleware/auth"
	"goiris/admin/app/middleware/jwt"
	"goiris/admin/app/middleware/sockets"
	"goiris/admin/app/web/router"
	"goiris/common"
	"goiris/common/storage"
	"goiris/common/support"
)

/**
 * 项目地址：https://gitee.com/yhm_my/go-iris
 */
func main() {
	common.InitConfig(conf.Asset)
	support.InitLog()
	support.InitValidator()
	storage.InitGorm()
	storage.InitRedis()
	auth.InitAuth()
	jwt.InitJWT()

	app := iris.New()
	// Websocket
	sockets.InitWebsocket(app)
	// HTTP
	router.Hub(app)
	if !common.G_AppConfig.Separate {
		app.RegisterView(iris.HTML("./resources", ".html").Binary(static.Asset, static.AssetNames))
		app.HandleDir("/", "./resources", iris.DirOptions{
			//IndexName:  "/index.html", // default "/index.html"
			Asset:      static.Asset,
			AssetNames: static.AssetNames,
			AssetInfo:  static.AssetInfo,
		})
	}
	// Start web server
	if err := app.Run(
		iris.Addr(fmt.Sprintf(":%d", common.G_AppConfig.Port)),
		iris.WithConfiguration(common.G_AppConfig.Configuration)); err != nil {
		golog.Fatalf("Start admin server failed. And err:%s", err.Error())
	}
}
