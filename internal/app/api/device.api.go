package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"dishes-admin-mod/internal/app/ginx"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/service"
)

var DeviceSet = wire.NewSet(wire.Struct(new(DeviceAPI), "*"))

type DeviceAPI struct {
	DeviceSrv *service.DeviceSrv
}

func (a *DeviceAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.DeviceQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.DeviceSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *DeviceAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.DeviceSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *DeviceAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Device
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.DeviceSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *DeviceAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Device
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.DeviceSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *DeviceAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.DeviceSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
