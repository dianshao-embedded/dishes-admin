//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"dishes-admin-mod/internal/app/api"
	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/module/adapter"
	"dishes-admin-mod/internal/app/router"
	"dishes-admin-mod/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitFileServer,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		InitUpdate,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
