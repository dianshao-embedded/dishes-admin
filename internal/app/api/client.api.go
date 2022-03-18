package api

import (
	"dishes-admin-mod/internal/app/ginx"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ClientSet = wire.NewSet(wire.Struct(new(ClientAPI), "*"))

type ClientAPI struct {
	ClientSrv *service.ClientSrv
}

func (a *ClientAPI) UpdateCommand(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := a.ClientSrv.UpdateCommand(ctx, ginx.ParseParamID(c, "device_id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *ClientAPI) UpdateEvent(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Client
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.ClientSrv.UpdateEvent(ctx, ginx.ParseParamID(c, "upgrade_id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	ginx.ResOK(c)
}

func (a *ClientAPI) WireguardEvent(c *gin.Context) {

}

func (a *ClientAPI) WireguardCommand(c *gin.Context) {

}
