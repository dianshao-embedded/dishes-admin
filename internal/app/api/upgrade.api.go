package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"dishes-admin-mod/internal/app/ginx"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/service"
)

var UpgradeSet = wire.NewSet(wire.Struct(new(UpgradeAPI), "*"))

type UpgradeAPI struct {
	UpgradeSrv *service.UpgradeSrv
}

func (a *UpgradeAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.UpgradeQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.UpgradeSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *UpgradeAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.UpgradeSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *UpgradeAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Upgrade
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.UpgradeSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *UpgradeAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Upgrade
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.UpgradeSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *UpgradeAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.UpgradeSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
