package mapper

import (
	"bytes"
	"fmt"
	"github.com/kataras/golog"
	"goiris/admin/app/web/vo"
	"goiris/common"
	"goiris/common/model"
	"goiris/common/storage"
)

type RoleMapper struct {}

func (rm *RoleMapper) Insert(vo *vo.AcceptRoleVO) error {
	return rm.createOrUpdate(vo, true)
}

func (rm *RoleMapper) FindOne(cond *model.CasbinRule) (*model.CasbinRule, error) {
	var (
		err   error
		resut = model.CasbinRule{}
	)
	if err = storage.G_DB.Where(&cond).First(&resut).Error; err != nil {
		return nil, err
	}
	return &resut, nil
}

func (rm *RoleMapper) UpdateOne(vo *vo.AcceptRoleVO) error {
	return rm.createOrUpdate(vo, false)
}

func (rm *RoleMapper) Delete(roleKeys []string) error {
	var (
		err error
		tx = storage.G_DB.Begin()
	)
	if err = tx.Delete(model.CasbinRule{}, "p_type='g' AND v1 in (?) AND v2=?", roleKeys, common.G_AppConfig.Domain).Error; err != nil {
		goto ERR
	}
	if err = tx.Delete(model.RoleMenu{}, "role_key in (?)", roleKeys).Error; err != nil {
		goto ERR
	}
	tx.Commit()
	return nil
ERR:
	tx.Rollback()
	return err
}

func (rm *RoleMapper) DomainList() (result []string, err error) {
	rows, err := storage.G_DB.Raw("SELECT v2 FROM casbin_rule WHERE p_type='g' GROUP BY v2").Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a string
		rows.Scan(&a)
		result = append(result, a)
	}
	return
}

func (rm *RoleMapper) Group() ([]*model.CasbinRule, error) {
	var (
		err    error
		result = make([]*model.CasbinRule, 0)
	)
	if err = storage.G_DB.Group("v1").Find(&result, "p_type='g' AND v2=?", common.G_AppConfig.Domain).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (rm *RoleMapper) Table(vo *vo.RoleVO) ([]*model.CasbinRule, error) {
	var (
		err    error
		menuList []*model.RoleMenu
		policyList []*model.CasbinRule
		result = make([]*model.CasbinRule, 0)
		tx = storage.G_DB.Begin()
	)
	if vo.Domain == "" || vo.Domain == "all" {
		if err = tx.Group("v1").Find(&result, "p_type='g'").Error; err != nil {
			goto ERR
		}
	} else {
		if err = tx.Group("v1").Find(&result, "p_type='g' AND v2=?", vo.Domain).Error; err != nil {
			goto ERR
		}
	}
	//
	for _, v := range result {
		// 角色对应的菜单集合
		if err = tx.Table("role_menu").Find(&menuList, "role_key=?", v.V1).Error; err != nil {
			goto ERR
		}
		v.MenuList = menuList
		// 角色对应的权限集合
		if err = tx.Find(&policyList, "p_type='p' AND v0=? AND v1=?", v.V1, v.V2).Error; err != nil {
			goto ERR
		}
		v.PolicyList = policyList
	}
	tx.Commit()
	return result, nil
ERR:
	tx.Rollback()
	return nil, err
}

func (rm *RoleMapper) RoleOfMenus(roleKey string) ([]*model.RoleMenu, error) {
	var (
		err    error
		result []*model.RoleMenu
	)
	if err = storage.G_DB.Table("role_menu").Find(&result, "role_key=?", roleKey).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (rm *RoleMapper) RoleOfPolicys(roleKey, domain string) ([]*model.CasbinRule, error) {
	var (
		err    error
		result []*model.CasbinRule
	)
	if err = storage.G_DB.Find(&result, "p_type='p' AND v0=? AND v1=?", roleKey, domain).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (rm *RoleMapper) PolicyTable(vo *vo.RoleVO) ([]*model.CasbinRule, error) {
	var (
		err    error
		result = make([]*model.CasbinRule, 0)
	)
	if vo.Domain == "" || vo.Domain == "all" {
		if err = storage.G_DB.Find(&result, "p_type='p'").Error; err != nil {
			goto ERR
		}
	} else {
		if err = storage.G_DB.Find(&result, "p_type='p' AND v1=?", vo.Domain).Error; err != nil {
			goto ERR
		}
	}
	return result, nil
ERR:
	return nil, err
}

// 更改角色表,角色-菜单表
func (rm *RoleMapper) createOrUpdate(vo *vo.AcceptRoleVO, isCreate bool) error {
	var (
		err error
		tx = storage.G_DB.Begin()
		roleKey = vo.Role.V1
		buffer bytes.Buffer
		menuSql = "insert into `role_menu` (`role_key`,`mid`) values"
		//policySql = "insert into `casbin_rule` (`p_type`,`v0`,`v1`,`v2`,`v5`) values"
	)
	// 角色表
	if isCreate {
		if err = tx.Model(&model.CasbinRule{}).Create(&vo.Role).Error; err != nil {
			goto ERR
		}
		roleKey = vo.Role.V1
	} else {
		if err = storage.G_DB.Model(&model.CasbinRule{}).Update(&vo.Role).Error; err != nil {
			goto ERR
		}
	}
	/* 角色 */
	// 删除的
	if vo.DeleteMids != nil && len(vo.DeleteMids) > 0 {
		if err = tx.Delete(&model.RoleMenu{}, "role_key=? AND mid in (?)", roleKey, vo.DeleteMids).Error; err != nil {
			goto ERR
		}
	}
	// 添加的
	if vo.AddMids != nil && len(vo.AddMids) > 0 {
		if _, err = buffer.WriteString(menuSql); err != nil {
			goto ERR
		}
		for i, v := range vo.AddMids {
			if i == len(vo.AddMids)-1 {
				buffer.WriteString(fmt.Sprintf("('%s',%d);", roleKey, v))
			} else {
				buffer.WriteString(fmt.Sprintf("('%s',%d),", roleKey, v))
			}
		}
		if err = tx.Exec(buffer.String()).Error; err != nil {
			golog.Errorf("更新角色菜单表信息失败。错误：%s", err)
		}
	}
	tx.Commit()
	return nil
ERR:
	tx.Rollback()
	return err
}
