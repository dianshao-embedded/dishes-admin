package service

import (
	"context"

	"github.com/google/wire"

	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
	"dishes-admin-mod/pkg/util/snowflake"
)

var UpgradeSet = wire.NewSet(wire.Struct(new(UpgradeSrv), "*"))

type UpgradeSrv struct {
	TransRepo   *dao.TransRepo
	UpgradeRepo *dao.UpgradeRepo
}

func (a *UpgradeSrv) Query(ctx context.Context, params schema.UpgradeQueryParam, opts ...schema.UpgradeQueryOptions) (*schema.UpgradeQueryResult, error) {
	result, err := a.UpgradeRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *UpgradeSrv) Get(ctx context.Context, id uint64, opts ...schema.UpgradeQueryOptions) (*schema.Upgrade, error) {
	item, err := a.UpgradeRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *UpgradeSrv) Create(ctx context.Context, item schema.Upgrade) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.UpgradeRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *UpgradeSrv) Update(ctx context.Context, id uint64, item schema.Upgrade) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.UpgradeRepo.Update(ctx, id, item)
	})
}

func (a *UpgradeSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.UpgradeRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.UpgradeRepo.Delete(ctx, id)
	})
}
