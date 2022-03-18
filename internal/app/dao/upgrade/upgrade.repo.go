package upgrade

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
)

// Injection wire
var UpgradeSet = wire.NewSet(wire.Struct(new(UpgradeRepo), "*"))

type UpgradeRepo struct {
	DB *gorm.DB
}

func (a *UpgradeRepo) getQueryOption(opts ...schema.UpgradeQueryOptions) schema.UpgradeQueryOptions {
	var opt schema.UpgradeQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *UpgradeRepo) Query(ctx context.Context, params schema.UpgradeQueryParam, opts ...schema.UpgradeQueryOptions) (*schema.UpgradeQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetUpgradeDB(ctx, a.DB)

	// TODO: Your where condition code here...
	if v := params.DeviceID; v != 0 {
		db = db.Where("device_id=?", v)
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

	var list Upgrades
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.UpgradeQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaUpgrades(),
	}

	return qr, nil
}

func (a *UpgradeRepo) Get(ctx context.Context, id uint64, opts ...schema.UpgradeQueryOptions) (*schema.Upgrade, error) {
	var item Upgrade
	ok, err := util.FindOne(ctx, GetUpgradeDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaUpgrade(), nil
}

func (a *UpgradeRepo) Create(ctx context.Context, item schema.Upgrade) error {
	eitem := SchemaUpgrade(item).ToUpgrade()
	result := GetUpgradeDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *UpgradeRepo) Update(ctx context.Context, id uint64, item schema.Upgrade) error {
	eitem := SchemaUpgrade(item).ToUpgrade()
	result := GetUpgradeDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *UpgradeRepo) Delete(ctx context.Context, id uint64) error {
	result := GetUpgradeDB(ctx, a.DB).Where("id=?", id).Delete(Upgrade{})
	return errors.WithStack(result.Error)
}
