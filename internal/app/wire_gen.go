// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"dishes-admin-mod/internal/app/api"
	"dishes-admin-mod/internal/app/dao/demo"
	"dishes-admin-mod/internal/app/dao/device"
	"dishes-admin-mod/internal/app/dao/firmware"
	"dishes-admin-mod/internal/app/dao/menu"
	"dishes-admin-mod/internal/app/dao/product"
	"dishes-admin-mod/internal/app/dao/remote"
	"dishes-admin-mod/internal/app/dao/role"
	"dishes-admin-mod/internal/app/dao/upgrade"
	"dishes-admin-mod/internal/app/dao/user"
	"dishes-admin-mod/internal/app/dao/util"
	"dishes-admin-mod/internal/app/module/adapter"
	"dishes-admin-mod/internal/app/router"
	"dishes-admin-mod/internal/app/service"
)

import (
	_ "dishes-admin-mod/internal/app/swagger"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	roleRepo := &role.RoleRepo{
		DB: db,
	}
	roleMenuRepo := &role.RoleMenuRepo{
		DB: db,
	}
	menuActionResourceRepo := &menu.MenuActionResourceRepo{
		DB: db,
	}
	userRepo := &user.UserRepo{
		DB: db,
	}
	userRoleRepo := &user.UserRoleRepo{
		DB: db,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleRepo:         roleRepo,
		RoleMenuRepo:     roleMenuRepo,
		MenuResourceRepo: menuActionResourceRepo,
		UserRepo:         userRepo,
		UserRoleRepo:     userRoleRepo,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	menuRepo := &menu.MenuRepo{
		DB: db,
	}
	menuActionRepo := &menu.MenuActionRepo{
		DB: db,
	}
	loginSrv := &service.LoginSrv{
		Auth:           auther,
		UserRepo:       userRepo,
		UserRoleRepo:   userRoleRepo,
		RoleRepo:       roleRepo,
		RoleMenuRepo:   roleMenuRepo,
		MenuRepo:       menuRepo,
		MenuActionRepo: menuActionRepo,
	}
	loginAPI := &api.LoginAPI{
		LoginSrv: loginSrv,
	}
	trans := &util.Trans{
		DB: db,
	}
	menuSrv := &service.MenuSrv{
		TransRepo:              trans,
		MenuRepo:               menuRepo,
		MenuActionRepo:         menuActionRepo,
		MenuActionResourceRepo: menuActionResourceRepo,
	}
	menuAPI := &api.MenuAPI{
		MenuSrv: menuSrv,
	}
	roleSrv := &service.RoleSrv{
		Enforcer:               syncedEnforcer,
		TransRepo:              trans,
		RoleRepo:               roleRepo,
		RoleMenuRepo:           roleMenuRepo,
		UserRepo:               userRepo,
		MenuActionResourceRepo: menuActionResourceRepo,
	}
	roleAPI := &api.RoleAPI{
		RoleSrv: roleSrv,
	}
	userSrv := &service.UserSrv{
		Enforcer:     syncedEnforcer,
		TransRepo:    trans,
		UserRepo:     userRepo,
		UserRoleRepo: userRoleRepo,
		RoleRepo:     roleRepo,
	}
	userAPI := &api.UserAPI{
		UserSrv: userSrv,
	}
	demoRepo := &demo.DemoRepo{
		DB: db,
	}
	demoSrv := &service.DemoSrv{
		TransRepo: trans,
		DemoRepo:  demoRepo,
	}
	demoAPI := &api.DemoAPI{
		DemoSrv: demoSrv,
	}
	productRepo := &product.ProductRepo{
		DB: db,
	}
	productSrv := &service.ProductSrv{
		TransRepo:   trans,
		ProductRepo: productRepo,
	}
	productAPI := &api.ProductAPI{
		ProductSrv: productSrv,
	}
	unroutedHandler, cleanup4, err := InitFileServer()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	firmwareRepo := &firmware.FirmwareRepo{
		DB: db,
	}
	deviceRepo := &device.DeviceRepo{
		DB: db,
	}
	upgradeRepo := &upgrade.UpgradeRepo{
		DB: db,
	}
	firmwareSrv := &service.FirmwareSrv{
		TransRepo:    trans,
		FirmwareRepo: firmwareRepo,
		DeviceRepo:   deviceRepo,
		UpgradeRepo:  upgradeRepo,
	}
	firmwareAPI := &api.FirmwareAPI{
		FirmwareSrv: firmwareSrv,
	}
	update, cleanup5, err := InitUpdate()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientSrv := &service.ClientSrv{
		DeviceRepo:  deviceRepo,
		UpgradeRepo: upgradeRepo,
		Update:      update,
	}
	clientAPI := &api.ClientAPI{
		ClientSrv: clientSrv,
	}
	deviceSrv := &service.DeviceSrv{
		TransRepo:  trans,
		DeviceRepo: deviceRepo,
	}
	deviceAPI := &api.DeviceAPI{
		DeviceSrv: deviceSrv,
	}
	upgradeSrv := &service.UpgradeSrv{
		TransRepo:   trans,
		UpgradeRepo: upgradeRepo,
	}
	upgradeAPI := &api.UpgradeAPI{
		UpgradeSrv: upgradeSrv,
	}
	remoteRepo := &remote.RemoteRepo{
		DB: db,
	}
	remoteSrv := &service.RemoteSrv{
		TransRepo:  trans,
		RemoteRepo: remoteRepo,
	}
	remoteAPI := &api.RemoteAPI{
		RemoteSrv: remoteSrv,
	}
	routerRouter := &router.Router{
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		LoginAPI:       loginAPI,
		MenuAPI:        menuAPI,
		RoleAPI:        roleAPI,
		UserAPI:        userAPI,
		DemoAPI:        demoAPI,
		ProductAPI:     productAPI,
		Handler:        unroutedHandler,
		FirmwareAPI:    firmwareAPI,
		ClientAPI:      clientAPI,
		DeviceAPI:      deviceAPI,
		UpgradeAPI:     upgradeAPI,
		RemoteAPI:      remoteAPI,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuSrv:        menuSrv,
		RollSrv:        roleSrv,
		UserSrv:        userSrv,
	}
	return injector, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}