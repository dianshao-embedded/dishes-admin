package service

import (
	"context"

	"github.com/google/wire"

	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
	"dishes-admin-mod/pkg/util/snowflake"
)

var RemoteSet = wire.NewSet(wire.Struct(new(RemoteSrv), "*"))

type RemoteSrv struct {
	TransRepo  *dao.TransRepo
	RemoteRepo *dao.RemoteRepo
}

func (a *RemoteSrv) Query(ctx context.Context, params schema.RemoteQueryParam, opts ...schema.RemoteQueryOptions) (*schema.RemoteQueryResult, error) {
	result, err := a.RemoteRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *RemoteSrv) Get(ctx context.Context, id uint64, opts ...schema.RemoteQueryOptions) (*schema.Remote, error) {
	item, err := a.RemoteRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *RemoteSrv) Create(ctx context.Context, item schema.Remote) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RemoteRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *RemoteSrv) Update(ctx context.Context, id uint64, item schema.Remote) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RemoteRepo.Update(ctx, id, item)
	})
}

func (a *RemoteSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.RemoteRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.RemoteRepo.Delete(ctx, id)
	})
}
