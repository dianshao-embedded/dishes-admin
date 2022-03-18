package api

import "github.com/google/wire"

var APISet = wire.NewSet(
	LoginSet,
	MenuSet,
	RoleSet,
	UserSet,
	DemoSet,
	ProductSet,
	FirmwareSet,
	ClientSet,
	DeviceSet,
	UpgradeSet,
	RemoteSet,
) // end
