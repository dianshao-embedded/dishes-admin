package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"dishes-admin-mod/internal/app/ginx"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/service"
)

var FirmwareSet = wire.NewSet(wire.Struct(new(FirmwareAPI), "*"))

type FirmwareAPI struct {
	FirmwareSrv *service.FirmwareSrv
}

func (a *FirmwareAPI) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.FirmwareQueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	result, err := a.FirmwareSrv.Query(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, result.Data, result.PageResult)
}

func (a *FirmwareAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.FirmwareSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item)
}

func (a *FirmwareAPI) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Firmware
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	result, err := a.FirmwareSrv.Create(ctx, item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, result)
}

func (a *FirmwareAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.Firmware
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResError(c, err)
		return
	}

	err := a.FirmwareSrv.Update(ctx, ginx.ParseParamID(c, "id"), item)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *FirmwareAPI) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.FirmwareSrv.Delete(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

func (a *FirmwareAPI) UploadFile(c *gin.Context) {
	ctx := c.Request.Context()
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	filename := header.Filename
	err = a.FirmwareSrv.UploadFile(ctx, file, filename, ginx.ParseParamID(c, "productID"), c.Param("stage"), c.Param("version"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}
