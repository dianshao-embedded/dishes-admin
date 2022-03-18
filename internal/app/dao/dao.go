package dao

import (
	"strings"

	"github.com/google/wire"
	"gorm.io/gorm"

	"dishes-admin-mod/internal/app/config"
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
) // end

// RepoSet repo injection
var RepoSet = wire.NewSet(
	util.TransSet,
	menu.MenuActionResourceSet,
	menu.MenuActionSet,
	menu.MenuSet,
	role.RoleMenuSet,
	role.RoleSet,
	user.UserRoleSet,
	user.UserSet,
	demo.DemoSet,
	product.ProductSet,
	firmware.FirmwareSet,
	device.DeviceSet,
	upgrade.UpgradeSet,
	remote.RemoteSet,
) // end

// Define repo type alias
type (
	TransRepo              = util.Trans
	MenuActionResourceRepo = menu.MenuActionResourceRepo
	MenuActionRepo         = menu.MenuActionRepo
	MenuRepo               = menu.MenuRepo
	RoleMenuRepo           = role.RoleMenuRepo
	RoleRepo               = role.RoleRepo
	UserRoleRepo           = user.UserRoleRepo
	UserRepo               = user.UserRepo
	DemoRepo               = demo.DemoRepo
	ProductRepo            = product.ProductRepo
	FirmwareRepo           = firmware.FirmwareRepo
	DeviceRepo             = device.DeviceRepo
	UpgradeRepo            = upgrade.UpgradeRepo
	RemoteRepo             = remote.RemoteRepo
) // end

// Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(menu.MenuActionResource),
		new(menu.MenuAction),
		new(menu.Menu),
		new(role.RoleMenu),
		new(role.Role),
		new(user.UserRole),
		new(user.User),
		new(demo.Demo),
		new(product.Product),
		new(product.Product),
		new(firmware.Firmware),
		new(device.Device),
		new(upgrade.Upgrade),
		new(remote.Remote),
	) // end
}
