package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"dishes-admin-mod/internal/app/ginx"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/service"
)

var RemoteSet = wire.NewSet(wire.Struct(new(RemoteAPI), "*"))

type RemoteAPI struct {
	RemoteSrv *service.RemoteSrv
}

func (a *RemoteAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.RemoteQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.RemoteSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *RemoteAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.RemoteSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *RemoteAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Remote
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.RemoteSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *RemoteAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Remote
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.RemoteSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *RemoteAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.RemoteSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
