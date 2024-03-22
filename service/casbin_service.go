package service

import (
	"memory/db"
	"memory/util"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var CasbinSrv *CasbinService

type CasbinService struct {
	Enforcer *casbin.Enforcer
	Adapter  *gormadapter.Adapter
}

// 初始化cabin服务
func InitCasbinService() {
	if CasbinSrv == nil {
		CasbinSrv, _ = NewCasbinService()
	}
}

func NewCasbinService() (*CasbinService, error) {
	a, err := gormadapter.NewAdapterByDB(db.DB)
	if err != nil {
		return nil, err
	}
	dir, err := util.GetPwd()
	if err != nil {
		return nil, err
	}
	m, err := model.NewModelFromFile(dir + "/config/rbac_model.conf")
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, err
	}
	return &CasbinService{Adapter: a, Enforcer: e}, nil
}

// (RoleName, Url, Method) 对应于 `CasbinRule` 表中的 (v0, v1, v2)
type RolePolicy struct {
	RoleName string `gorm:"column:v0"`
	Url      string `gorm:"column:v1"`
	Method   string `gorm:"column:v2"`
}

// 获取所有角色组
func (c *CasbinService) GetRoles() []string {
	return c.Enforcer.GetAllRoles()
}

// 获取所有角色组权限
func (c *CasbinService) GetRolePolicy() (roles []RolePolicy, err error) {
	err = c.Adapter.GetDb().Model(&gormadapter.CasbinRule{}).Where("ptype = 'p'").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return
}

// 创建角色组权限, 已有的会忽略
func (c *CasbinService) CreateRolePolicy(r RolePolicy) error {
	// 不直接操作数据库，利用Enforcer简化操作
	err := c.Enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	_, err = c.Enforcer.AddPolicy(r.RoleName, r.Url, r.Method)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// 修改角色组权限
func (c *CasbinService) UpdateRolePolicy(old, new RolePolicy) error {
	_, err := c.Enforcer.UpdatePolicy([]string{old.RoleName, old.Url, old.Method},
		[]string{new.RoleName, new.Url, new.Method})
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// 删除角色组权限
func (c *CasbinService) DeleteRolePolicy(r RolePolicy) error {
	_, err := c.Enforcer.RemovePolicy(r.RoleName, r.Url, r.Method)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

type User struct {
	UserName  string
	RoleNames []string
}

// 获取所有用户以及关联的角色
func (c *CasbinService) GetUsers() (users []User) {
	p := c.Enforcer.GetGroupingPolicy()
	usernameUser := make(map[string]*User, 0)
	for _, _p := range p {
		username, usergroup := _p[0], _p[1]
		if v, ok := usernameUser[username]; ok {
			usernameUser[username].RoleNames = append(v.RoleNames, usergroup)
		} else {
			usernameUser[username] = &User{UserName: username, RoleNames: []string{usergroup}}
		}
	}
	for _, v := range usernameUser {
		users = append(users, *v)
	}
	return
}

// 角色组中添加用户, 没有组默认创建
func (c *CasbinService) UpdateUserRole(username, rolename string) error {
	_, err := c.Enforcer.AddGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// 角色组中删除用户
func (c *CasbinService) DeleteUserRole(username, rolename string) error {
	_, err := c.Enforcer.RemoveGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// 验证用户权限
func (c *CasbinService) CanAccess(username, url, method string) (ok bool, err error) {
	return c.Enforcer.Enforce(username, url, method)
}
