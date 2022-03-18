package service

import (
	"context"

	"github.com/google/wire"

	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
	"dishes-admin-mod/pkg/util/snowflake"
)

var ProductSet = wire.NewSet(wire.Struct(new(ProductSrv), "*"))

type ProductSrv struct {
	TransRepo   *dao.TransRepo
	ProductRepo *dao.ProductRepo
}

func (a *ProductSrv) Query(ctx context.Context, params schema.ProductQueryParam, opts ...schema.ProductQueryOptions) (*schema.ProductQueryResult, error) {
	result, err := a.ProductRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *ProductSrv) Get(ctx context.Context, id uint64, opts ...schema.ProductQueryOptions) (*schema.Product, error) {
	item, err := a.ProductRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *ProductSrv) Create(ctx context.Context, item schema.Product) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.ProductRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *ProductSrv) Update(ctx context.Context, id uint64, item schema.Product) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.ProductRepo.Update(ctx, id, item)
	})
}

func (a *ProductSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.ProductRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.ProductRepo.Delete(ctx, id)
	})
}
