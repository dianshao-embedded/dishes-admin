package service

import (
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	MenuSet,
	RoleSet,
	UserSet,
	LoginSet,
	DemoSet,
	ProductSet,
	FirmwareSet,
	ClientSet,
	DeviceSet,
	UpgradeSet,
	RemoteSet,
) // end
