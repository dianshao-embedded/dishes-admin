package product

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
)

// Injection wire
var ProductSet = wire.NewSet(wire.Struct(new(ProductRepo), "*"))

type ProductRepo struct {
	DB *gorm.DB
}

func (a *ProductRepo) getQueryOption(opts ...schema.ProductQueryOptions) schema.ProductQueryOptions {
	var opt schema.ProductQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *ProductRepo) Query(ctx context.Context, params schema.ProductQueryParam, opts ...schema.ProductQueryOptions) (*schema.ProductQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetProductDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Products
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.ProductQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaProducts(),
	}

	return qr, nil
}

func (a *ProductRepo) Get(ctx context.Context, id uint64, opts ...schema.ProductQueryOptions) (*schema.Product, error) {
	var item Product
	ok, err := util.FindOne(ctx, GetProductDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaProduct(), nil
}

func (a *ProductRepo) Create(ctx context.Context, item schema.Product) error {
	eitem := SchemaProduct(item).ToProduct()
	result := GetProductDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *ProductRepo) Update(ctx context.Context, id uint64, item schema.Product) error {
	eitem := SchemaProduct(item).ToProduct()
	result := GetProductDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *ProductRepo) Delete(ctx context.Context, id uint64) error {
	result := GetProductDB(ctx, a.DB).Where("id=?", id).Delete(Product{})
	return errors.WithStack(result.Error)
}
