package service

import (
	"context"

	"github.com/google/wire"

	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
	"dishes-admin-mod/pkg/util/snowflake"
)

var DeviceSet = wire.NewSet(wire.Struct(new(DeviceSrv), "*"))

type DeviceSrv struct {
	TransRepo  *dao.TransRepo
	DeviceRepo *dao.DeviceRepo
}

func (a *DeviceSrv) Query(ctx context.Context, params schema.DeviceQueryParam, opts ...schema.DeviceQueryOptions) (*schema.DeviceQueryResult, error) {
	result, err := a.DeviceRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *DeviceSrv) Get(ctx context.Context, id uint64, opts ...schema.DeviceQueryOptions) (*schema.Device, error) {
	item, err := a.DeviceRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *DeviceSrv) Create(ctx context.Context, item schema.Device) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DeviceRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *DeviceSrv) Update(ctx context.Context, id uint64, item schema.Device) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DeviceRepo.Update(ctx, id, item)
	})
}

func (a *DeviceSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.DeviceRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.DeviceRepo.Delete(ctx, id)
	})
}
