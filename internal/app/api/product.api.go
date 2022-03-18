package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"dishes-admin-mod/internal/app/ginx"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/service"
)

var ProductSet = wire.NewSet(wire.Struct(new(ProductAPI), "*"))

type ProductAPI struct {
	ProductSrv *service.ProductSrv
}

func (a *ProductAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.ProductQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.ProductSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *ProductAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.ProductSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *ProductAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Product
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.ProductSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *ProductAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Product
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.ProductSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *ProductAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.ProductSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
