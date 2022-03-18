package device

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
)

// Injection wire
var DeviceSet = wire.NewSet(wire.Struct(new(DeviceRepo), "*"))

type DeviceRepo struct {
	DB *gorm.DB
}

func (a *DeviceRepo) getQueryOption(opts ...schema.DeviceQueryOptions) schema.DeviceQueryOptions {
	var opt schema.DeviceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *DeviceRepo) Query(ctx context.Context, params schema.DeviceQueryParam, opts ...schema.DeviceQueryOptions) (*schema.DeviceQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetDeviceDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.ProductID; v != 0 {
		db = db.Where("product_id=?", v)
	}

	if v := params.Name; v != "" {
		db = db.Where("name=?", v)
	}

	if v := params.SN; v != "" {
		db = db.Where("sn=?", v)
	}

	if v := params.Stage; v != "" {
		db = db.Where("stage=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Devices
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.DeviceQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaDevices(),
	}

	return qr, nil
}

func (a *DeviceRepo) Get(ctx context.Context, id uint64, opts ...schema.DeviceQueryOptions) (*schema.Device, error) {
	var item Device
	ok, err := util.FindOne(ctx, GetDeviceDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDevice(), nil
}

func (a *DeviceRepo) Create(ctx context.Context, item schema.Device) error {
	eitem := SchemaDevice(item).ToDevice()
	result := GetDeviceDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *DeviceRepo) Update(ctx context.Context, id uint64, item schema.Device) error {
	eitem := SchemaDevice(item).ToDevice()
	result := GetDeviceDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *DeviceRepo) Delete(ctx context.Context, id uint64) error {
	result := GetDeviceDB(ctx, a.DB).Where("id=?", id).Delete(Device{})
	return errors.WithStack(result.Error)
}
