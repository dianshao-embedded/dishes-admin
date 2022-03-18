package demo

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
)

// Injection wire
var DemoSet = wire.NewSet(wire.Struct(new(DemoRepo), "*"))

type DemoRepo struct {
	DB *gorm.DB
}

func (a *DemoRepo) getQueryOption(opts ...schema.DemoQueryOptions) schema.DemoQueryOptions {
	var opt schema.DemoQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *DemoRepo) Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetDemoDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Demos
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.DemoQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaDemos(),
	}

	return qr, nil
}

func (a *DemoRepo) Get(ctx context.Context, id uint64, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {
	var item Demo
	ok, err := util.FindOne(ctx, GetDemoDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaDemo(), nil
}

func (a *DemoRepo) Create(ctx context.Context, item schema.Demo) error {
	eitem := SchemaDemo(item).ToDemo()
	result := GetDemoDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *DemoRepo) Update(ctx context.Context, id uint64, item schema.Demo) error {
	eitem := SchemaDemo(item).ToDemo()
	result := GetDemoDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *DemoRepo) Delete(ctx context.Context, id uint64) error {
	result := GetDemoDB(ctx, a.DB).Where("id=?", id).Delete(Demo{})
	return errors.WithStack(result.Error)
}
