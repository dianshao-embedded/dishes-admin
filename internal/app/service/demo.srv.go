package service

import (
	"context"

	"github.com/google/wire"

	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
	"dishes-admin-mod/pkg/util/snowflake"
)

var DemoSet = wire.NewSet(wire.Struct(new(DemoSrv), "*"))

type DemoSrv struct {
	TransRepo *dao.TransRepo
	DemoRepo  *dao.DemoRepo
}

func (a *DemoSrv) Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error) {
	result, err := a.DemoRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *DemoSrv) Get(ctx context.Context, id uint64, opts ...schema.DemoQueryOptions) (*schema.Demo, error) {
	item, err := a.DemoRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *DemoSrv) Create(ctx context.Context, item schema.Demo) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DemoRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *DemoSrv) Update(ctx context.Context, id uint64, item schema.Demo) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DemoRepo.Update(ctx, id, item)
	})
}

func (a *DemoSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.DemoRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DemoRepo.Delete(ctx, id)
	})
}
