package service

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/google/wire"

	"dishes-admin-mod/internal/app/config"
	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/pkg/errors"
	"dishes-admin-mod/pkg/util/snowflake"
)

var FirmwareSet = wire.NewSet(wire.Struct(new(FirmwareSrv), "*"))

type FirmwareSrv struct {
	TransRepo    *dao.TransRepo
	FirmwareRepo *dao.FirmwareRepo
	DeviceRepo   *dao.DeviceRepo
	UpgradeRepo  *dao.UpgradeRepo
}

func (a *FirmwareSrv) Query(ctx context.Context, params schema.FirmwareQueryParam, opts ...schema.FirmwareQueryOptions) (*schema.FirmwareQueryResult, error) {
	result, err := a.FirmwareRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *FirmwareSrv) Get(ctx context.Context, id uint64, opts ...schema.FirmwareQueryOptions) (*schema.Firmware, error) {
	item, err := a.FirmwareRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *FirmwareSrv) Create(ctx context.Context, item schema.Firmware) (*schema.IDResult, error) {
	item.ID = snowflake.MustID()

	err := a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		// 将固件更新信息同步到产品下的相应阶段设备

		result, err := a.DeviceRepo.Query(ctx, schema.DeviceQueryParam{
			ProductID: item.ProductID,
			Stage:     item.Stage,
		})

		if err != nil {
			return err
		}

		for _, v := range result.Data {
			err = a.UpgradeRepo.Create(ctx, schema.Upgrade{
				ID:         snowflake.MustID(),
				FirmwareID: item.ID,
				DeviceID:   v.ID,
				Version:    item.Version,
				Status:     "pending",
				Failure:    0,
				Name:       item.Name,
				Size:       item.Size,
				Stage:      item.Stage,
			})

			if err != nil {
				return err
			}
		}

		return a.FirmwareRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *FirmwareSrv) Update(ctx context.Context, id uint64, item schema.Firmware) error {
	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	item.ID = oldItem.ID
	item.CreatedAt = oldItem.CreatedAt

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.FirmwareRepo.Update(ctx, id, item)
	})
}

func (a *FirmwareSrv) Delete(ctx context.Context, id uint64) error {
	//TODO: 删除固件怎么做，需要删除所有设备下的相关记录
	oldItem, err := a.FirmwareRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		return a.FirmwareRepo.Delete(ctx, id)
	})
}

func (a *FirmwareSrv) UploadFile(ctx context.Context, file multipart.File, filename string, productID uint64, stage string, version string) error {
	d := config.C.FileServer.Directory + "/" + strconv.FormatUint(productID, 10) + "/" + stage + "/" + version + "/"

	exist, err := a.pathExists(d)
	if err != nil {
		return err
	}
	if !exist {
		os.MkdirAll(d, os.ModePerm)
	}

	exist, err = a.pathExists(d + filename)
	if err != nil {
		return err
	}
	if exist {
		if stage == "dev" {
			os.Remove(d + filename)
		} else {
			return errors.New("firmware exists")
		}
	}

	out, err := os.Create(d + filename)
	if err != nil {
		return err
	}

	defer out.Close()
	println(1)
	_, err = io.Copy(out, file)
	println(2)
	if err != nil {
		return err
	}

	return nil
}

func (a *FirmwareSrv) pathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
