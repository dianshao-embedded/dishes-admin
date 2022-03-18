package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var FirmwareSet = wire.NewSet(wire.Struct(new(FirmwareMock), "*"))

type FirmwareMock struct{}

// @Tags 固件
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param productID query int true "产品ID"
// @Success 200 {object} schema.ListResult{list=[]schema.Firmware} "Response Data"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/firmwares [get]
func (a *FirmwareMock) Query(c *gin.Context) {
}

// @Tags 固件
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.Firmware
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/firmwares/{id} [get]
func (a *FirmwareMock) Get(c *gin.Context) {
}

// @Tags 固件
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param body body schema.Firmware true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/firmwares [post]
func (a *FirmwareMock) Create(c *gin.Context) {
}

// @Tags 固件
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Param body body schema.Firmware true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:bad request}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/firmwares/{id} [put]
func (a *FirmwareMock) Update(c *gin.Context) {
}

// @Tags 固件
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param id path int true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:9999,message:invalid signature}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:internal server error}}"
// @Router /api/v1/firmwares/{id} [delete]
func (a *FirmwareMock) Delete(c *gin.Context) {
}
