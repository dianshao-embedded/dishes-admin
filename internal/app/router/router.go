package router

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/tus/tusd/pkg/handler"

	"dishes-admin-mod/internal/app/api"
	"dishes-admin-mod/internal/app/config"
	"dishes-admin-mod/internal/app/middleware"
	"dishes-admin-mod/pkg/auth"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	Auth           auth.Auther
	CasbinEnforcer *casbin.SyncedEnforcer
	LoginAPI       *api.LoginAPI
	MenuAPI        *api.MenuAPI
	RoleAPI        *api.RoleAPI
	UserAPI        *api.UserAPI
	DemoAPI        *api.DemoAPI
	ProductAPI     *api.ProductAPI
	Handler        *handler.UnroutedHandler
	FirmwareAPI    *api.FirmwareAPI
	ClientAPI      *api.ClientAPI
	DeviceAPI      *api.DeviceAPI
	UpgradeAPI     *api.UpgradeAPI
	RemoteAPI      *api.RemoteAPI
} // end

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
	))

	g.Use(middleware.CasbinMiddleware(a.CasbinEnforcer,
		middleware.AllowPathPrefixSkipper("/api/v1/pub"),
	))

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{
		pub := v1.Group("/pub")
		{
			gLogin := pub.Group("login")
			{
				gLogin.GET("captchaid", a.LoginAPI.GetCaptcha)
				gLogin.GET("captcha", a.LoginAPI.ResCaptcha)
				//gLogin.POST("", a.LoginAPI.Login)
				gLogin.POST("", a.LoginAPI.LoginWithoutCaptcha)
				gLogin.POST("exit", a.LoginAPI.Logout)
			}

			gCurrent := pub.Group("current")
			{
				gCurrent.PUT("password", a.LoginAPI.UpdatePassword)
				gCurrent.GET("user", a.LoginAPI.GetUserInfo)
				gCurrent.GET("menutree", a.LoginAPI.QueryUserMenuTree)
			}
			pub.POST("/refresh-token", a.LoginAPI.RefreshToken)
		}

		gMenu := v1.Group("menus")
		{
			gMenu.GET("", a.MenuAPI.Query)
			gMenu.GET(":id", a.MenuAPI.Get)
			gMenu.POST("", a.MenuAPI.Create)
			gMenu.PUT(":id", a.MenuAPI.Update)
			gMenu.DELETE(":id", a.MenuAPI.Delete)
			gMenu.PATCH(":id/enable", a.MenuAPI.Enable)
			gMenu.PATCH(":id/disable", a.MenuAPI.Disable)
		}
		v1.GET("/menus.tree", a.MenuAPI.QueryTree)

		gRole := v1.Group("roles")
		{
			gRole.GET("", a.RoleAPI.Query)
			gRole.GET(":id", a.RoleAPI.Get)
			gRole.POST("", a.RoleAPI.Create)
			gRole.PUT(":id", a.RoleAPI.Update)
			gRole.DELETE(":id", a.RoleAPI.Delete)
			gRole.PATCH(":id/enable", a.RoleAPI.Enable)
			gRole.PATCH(":id/disable", a.RoleAPI.Disable)
		}
		v1.GET("/roles.select", a.RoleAPI.QuerySelect)

		gUser := v1.Group("users")
		{
			gUser.GET("", a.UserAPI.Query)
			gUser.GET(":id", a.UserAPI.Get)
			gUser.POST("", a.UserAPI.Create)
			gUser.PUT(":id", a.UserAPI.Update)
			gUser.DELETE(":id", a.UserAPI.Delete)
			gUser.PATCH(":id/enable", a.UserAPI.Enable)
			gUser.PATCH(":id/disable", a.UserAPI.Disable)
		}

		gDemo := v1.Group("demos")
		{
			gDemo.GET("", a.DemoAPI.Query)
			gDemo.GET(":id", a.DemoAPI.Get)
			gDemo.POST("", a.DemoAPI.Create)
			gDemo.PUT(":id", a.DemoAPI.Update)
			gDemo.DELETE(":id", a.DemoAPI.Delete)

		}

		gProduct := v1.Group("products")
		{
			gProduct.GET("", a.ProductAPI.Query)
			gProduct.GET(":id", a.ProductAPI.Get)
			gProduct.POST("", a.ProductAPI.Create)
			gProduct.PUT(":id", a.ProductAPI.Update)
			gProduct.DELETE(":id", a.ProductAPI.Delete)

		}

		gFirmware := v1.Group("firmwares")
		{
			gFirmware.GET("", a.FirmwareAPI.Query)
			gFirmware.GET(":id", a.FirmwareAPI.Get)
			gFirmware.POST("", a.FirmwareAPI.Create)
			gFirmware.POST("upload/:productID/:stage/:version", a.FirmwareAPI.UploadFile)
			gFirmware.PUT(":id", a.FirmwareAPI.Update)
			gFirmware.DELETE(":id", a.FirmwareAPI.Delete)
			gFirmware.StaticFS("downloads", http.Dir(config.C.FileServer.Directory))
		}

		gClient := v1.Group("client")
		{
			gClient.GET("update-command/:device_id", a.ClientAPI.UpdateCommand)
			gClient.POST("update-event/:upgrade_id", a.ClientAPI.UpdateEvent)
		}

		gDevice := v1.Group("devices")
		{
			gDevice.GET("", a.DeviceAPI.Query)
			gDevice.GET(":id", a.DeviceAPI.Get)
			gDevice.POST("", a.DeviceAPI.Create)
			gDevice.PUT(":id", a.DeviceAPI.Update)
			gDevice.DELETE(":id", a.DeviceAPI.Delete)
		}

		gUpgrade := v1.Group("upgrades")
		{
			gUpgrade.GET("", a.UpgradeAPI.Query)
			gUpgrade.GET(":id", a.UpgradeAPI.Get)
			gUpgrade.POST("", a.UpgradeAPI.Create)
			gUpgrade.PUT(":id", a.UpgradeAPI.Update)
			gUpgrade.DELETE(":id", a.UpgradeAPI.Delete)
		}

		gRemote := v1.Group("remotes")
		{
			gRemote.GET("", a.RemoteAPI.Query)
			gRemote.GET(":id", a.RemoteAPI.Get)
			gRemote.POST("", a.RemoteAPI.Create)
			gRemote.PUT(":id", a.RemoteAPI.Update)
			gRemote.DELETE(":id", a.RemoteAPI.Delete)

		}

	} // v1 end
}
