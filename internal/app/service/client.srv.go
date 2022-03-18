package service

import (
	"context"
	"dishes-admin-mod/internal/app/dao"
	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/update"
	"errors"
	"time"

	"github.com/google/wire"
)

var ClientSet = wire.NewSet(wire.Struct(new(ClientSrv), "*"))

//TODO: 多个产品同时更新
//TODO: 每个产品中包含多个终端，如何确定每个终端均更新成功

type ClientSrv struct {
	DeviceRepo  *dao.DeviceRepo
	UpgradeRepo *dao.UpgradeRepo
	Update      *update.Update
	RemoteRepo  *dao.RemoteRepo
}

func (a *ClientSrv) updateRoutine(ctx context.Context, upgrade_id uint64) {
	item, err := a.UpgradeRepo.Get(ctx, upgrade_id)
	if err != nil {
		return
	}

	item2, err := a.DeviceRepo.Get(ctx, item.DeviceID)
	if err != nil {
		return
	}

	item.Status = "upgrading"
	a.UpgradeRepo.Update(ctx, item.ID, *item)
	select {
	case <-a.Update.SuccessChan[upgrade_id]:
		item.Status = "success"
		a.UpgradeRepo.Update(ctx, item.ID, *item)
		item2.Version = item.Version
		a.DeviceRepo.Update(ctx, item2.ID, *item2)
		close(a.Update.FailedChan[upgrade_id])
	case <-a.Update.FailedChan[upgrade_id]:
		item.Failure++
		if item.Failure > 2 {
			item.Status = "failure"
		} else {
			item.Status = "pending"
		}
		a.UpgradeRepo.Update(ctx, item.ID, *item)
		close(a.Update.SuccessChan[upgrade_id])
	case <-time.After(time.Second * 600):
		item.Failure++
		if item.Failure > 2 {
			item.Status = "failure"
		} else {
			item.Status = "pending"
		}
		a.UpgradeRepo.Update(ctx, item.ID, *item)
		close(a.Update.SuccessChan[upgrade_id])
		close(a.Update.FailedChan[upgrade_id])
	}
}

func (a *ClientSrv) UpdateCommand(ctx context.Context, DeviceID uint64) (*schema.Client, error) {
	// TODO: 如果开始一个更新命令，需要启动一个更新线程，等待更新结果事件上报
	result, err := a.UpgradeRepo.Query(ctx, schema.UpgradeQueryParam{
		DeviceID: DeviceID,
		Status:   "pending",
	})
	if err != nil {
		return nil, errors.New("database query error")
	} else if len(result.Data) == 0 {
		return &schema.Client{
			Update: 0,
		}, nil
	} else if len(result.Data) == 1 {
		return &schema.Client{
			UpgradeID: result.Data[0].ID,
			Filename:  result.Data[0].Name,
			Verison:   result.Data[0].Version,
			Stage:     result.Data[0].Stage,
			Size:      result.Data[0].Size,
			Update:    1,
		}, nil
	} else {
		//TODO: 多个未更新固件应该采用何种策略
		return &schema.Client{}, nil
	}
}

func (a *ClientSrv) UpdateEvent(ctx context.Context, upgrade_id uint64, item schema.Client) error {
	switch item.Event {
	case "start_update":
		a.Update.SuccessChan[upgrade_id] = make(chan struct{}, 1)
		a.Update.FailedChan[upgrade_id] = make(chan struct{}, 1)
		go a.updateRoutine(ctx, upgrade_id)

	case "error_update":
		close(a.Update.FailedChan[upgrade_id])

	case "success_update":
		close(a.Update.SuccessChan[upgrade_id])
	}

	return nil
}

/*
func (a *ClientSrv) WireguardCommand() {

}

func (a *ClientSrv) WireguardEvent(ctx context.Context, item schema.Client, DeviceID uint64) error {
	result, err := a.RemoteRepo.Query(ctx, schema.RemoteQueryParam{})
	if err != nil {
		return err
	}

	ip_address := "172.20.1." + strconv.Itoa(len(result.Data)+2)

	err := a.RemoteRepo.Create(ctx, schema.Remote{
		DeviceID: DeviceID,
		Address:  item,
	})
}
*/
