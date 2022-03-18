package firmware

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
)

// Injection wire
var FirmwareSet = wire.NewSet(wire.Struct(new(FirmwareRepo), "*"))

type FirmwareRepo struct {
	DB *gorm.DB
}

func (a *FirmwareRepo) getQueryOption(opts ...schema.FirmwareQueryOptions) schema.FirmwareQueryOptions {
	var opt schema.FirmwareQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *FirmwareRepo) Query(ctx context.Context, params schema.FirmwareQueryParam, opts ...schema.FirmwareQueryOptions) (*schema.FirmwareQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetFirmwareDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.ProductID; v != 0 {
		db = db.Where("product_id=?", v)
	}

	if v := params.Status; v != "" {
		db = db.Where("status=?", v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Firmwares
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.FirmwareQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaFirmwares(),
	}

	return qr, nil
}

func (a *FirmwareRepo) Get(ctx context.Context, id uint64, opts ...schema.FirmwareQueryOptions) (*schema.Firmware, error) {
	var item Firmware
	ok, err := util.FindOne(ctx, GetFirmwareDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaFirmware(), nil
}

func (a *FirmwareRepo) Create(ctx context.Context, item schema.Firmware) error {
	eitem := SchemaFirmware(item).ToFirmware()
	result := GetFirmwareDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *FirmwareRepo) Update(ctx context.Context, id uint64, item schema.Firmware) error {
	eitem := SchemaFirmware(item).ToFirmware()
	result := GetFirmwareDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *FirmwareRepo) Delete(ctx context.Context, id uint64) error {
	result := GetFirmwareDB(ctx, a.DB).Where("id=?", id).Delete(Firmware{})
	return errors.WithStack(result.Error)
}
