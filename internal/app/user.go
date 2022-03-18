package app

import (
	"context"

	"dishes-admin-mod/internal/app/schema"
	"dishes-admin-mod/internal/app/service"
)

func InitRole(r *service.RoleSrv, m *service.MenuSrv) (*schema.IDResult, error) {
	ctx := context.Background()
	var params schema.MenuQueryParam
	params.Pagination = true
	mResult, err := m.Query(ctx, params, schema.MenuQueryOptions{
		OrderFields: schema.NewOrderFields(schema.NewOrderField("sequence", schema.OrderByDESC)),
	})
	if err != nil {
		return nil, err
	}

	mData := mResult.Data

	var params2 schema.RoleQueryParam
	params2.Name = "worker"
	rResult, err := r.Query(ctx, params2)
	if err != nil {
		return nil, err
	}
	if len(rResult.Data) == 0 {
		_, err = r.Create(ctx, schema.Role{
			Name:     "worker",
			Sequence: 1000000,
			Status:   1,
			RoleMenus: schema.RoleMenus{
				&schema.RoleMenu{
					MenuID:   mData[1].ID,
					ActionID: mData[1].Actions[0].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[1].ID,
					ActionID: mData[1].Actions[1].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[1].ID,
					ActionID: mData[1].Actions[2].ID,
				},
				&schema.RoleMenu{
					MenuID:   mData[1].ID,
					ActionID: mData[1].Actions[3].ID,
				},
			},
		})
		if err != nil {
			return nil, err
		}
	}

	params2.Name = "admin"
	rResult, err = r.Query(ctx, params2)
	if err != nil {
		return nil, err
	}

	var rm schema.RoleMenus

	for _, v1 := range mData {
		for _, v2 := range v1.Actions {
			rm = append(rm, &schema.RoleMenu{
				MenuID:   v1.ID,
				ActionID: v2.ID,
			})
		}
	}
	if len(rResult.Data) == 0 {
		return r.Create(ctx, schema.Role{
			Name:      "admin",
			Sequence:  1,
			Status:    1,
			RoleMenus: rm,
		})
	} else {
		return &schema.IDResult{ID: rResult.Data[0].ID}, nil
	}
}

func InitUser(u *service.UserSrv, r *service.RoleSrv, m *service.MenuSrv) error {
	ctx := context.Background()
	admin, err := InitRole(r, m)
	if err != nil {
		return err
	}

	param := schema.UserQueryParam{
		UserName: "admin",
	}
	result, err := u.Query(ctx, param)
	if err != nil {
		return err
	}

	if len(result.Data) > 0 {
		return nil
	} else {
		_, err = u.Create(ctx, schema.User{
			UserName: "admin",
			RealName: "admin",
			Password: "admin",
			Status:   1,
			UserRoles: schema.UserRoles{
				&schema.UserRole{
					RoleID: admin.ID,
				},
			},
		})
	}

	return err
}
