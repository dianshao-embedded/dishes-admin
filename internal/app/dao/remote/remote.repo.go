package remote

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
)

// Injection wire
var RemoteSet = wire.NewSet(wire.Struct(new(RemoteRepo), "*"))

type RemoteRepo struct {
	DB *gorm.DB
}

func (a *RemoteRepo) getQueryOption(opts ...schema.RemoteQueryOptions) schema.RemoteQueryOptions {
	var opt schema.RemoteQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *RemoteRepo) Query(ctx context.Context, params schema.RemoteQueryParam, opts ...schema.RemoteQueryOptions) (*schema.RemoteQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetRemoteDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Remotes
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.RemoteQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaRemotes(),
	}

	return qr, nil
}

func (a *RemoteRepo) Get(ctx context.Context, id uint64, opts ...schema.RemoteQueryOptions) (*schema.Remote, error) {
	var item Remote
	ok, err := util.FindOne(ctx, GetRemoteDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaRemote(), nil
}

func (a *RemoteRepo) Create(ctx context.Context, item schema.Remote) error {
	eitem := SchemaRemote(item).ToRemote()
	result := GetRemoteDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *RemoteRepo) Update(ctx context.Context, id uint64, item schema.Remote) error {
	eitem := SchemaRemote(item).ToRemote()
	result := GetRemoteDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *RemoteRepo) Delete(ctx context.Context, id uint64) error {
	result := GetRemoteDB(ctx, a.DB).Where("id=?", id).Delete(Remote{})
	return errors.WithStack(result.Error)
}
