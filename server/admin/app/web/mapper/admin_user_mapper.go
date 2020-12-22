package mapper

import (
	"bytes"
	"fmt"
	"github.com/kataras/golog"
	"goiris/admin/app/web/vo"
	"goiris/common"
	"goiris/common/model"
	"goiris/common/storage"
	"strconv"
)

type AdminUserMapper struct {}

func (aum *AdminUserMapper) Insert(vo *vo.AcceptAdminUserVO) error {
	return aum.createOrUpdate(vo, true)
}

func (aum *AdminUserMapper) FindOne(cond *model.AdminUser) (*model.AdminUser, error) {
	var (
		err   error
		resut = model.AdminUser{}
	)
	if err = storage.G_DB.Where(&cond).First(&resut).Error; err != nil {
		return nil, err
	}
	return &resut, nil
}

func (aum *AdminUserMapper) UpdateOne(vo *vo.AcceptAdminUserVO) error {
	return aum.createOrUpdate(vo, false)
}

func (aum *AdminUserMapper) Delete(ids []int64) error {
	var (
		err error
		tx = storage.G_DB.Begin()
		v0List []string
	)
	if err = tx.Delete(model.AdminUser{}, "id in (?)", ids).Error; err != nil {
		goto ERR
	}
	for _, v := range ids {
		v0List = append(v0List, strconv.FormatInt(v, 10))
	}
	if err = tx.Delete(model.CasbinRule{}, "p_type='g' AND v0 in (?) AND v2=?", v0List, common.G_AppConfig.Domain).Error; err != nil {
		goto ERR
	}
	tx.Commit()
	return nil
ERR:
	tx.Rollback()
	return err
}

func (aum *AdminUserMapper) FindList(vo *vo.AdminUserVO) (int64, []*model.AdminUser, error) {
	var (
		err    error
		total  int64
		result = make([]*model.AdminUser, 0)
	)
	if err = storage.G_DB.Model(&model.AdminUser{}).Count(&total).Error; err != nil {
		goto ERR
	}
	if err = storage.G_DB.
		Select("id, username, gender, enable, name, age, phone, email, userface, create_time, update_time, label").
		Offset(vo.Start).Limit(vo.Size).Find(&result).Error; err != nil {
		goto ERR
	}
	for _, v := range result {
		var roles []*model.CasbinRule
		if err = storage.G_DB.Where("v2 = 'a' AND p_type = 'g' AND v0 = ?", v.Id).Find(&roles).Error; err != nil {
			goto  ERR
		}
		v.Roles = roles
	}
	return total, result, nil
ERR:
	return 0, nil, err
}

// 更改用户表,角色表
func (aum *AdminUserMapper) createOrUpdate(vo *vo.AcceptAdminUserVO, isCreate bool) error {
	var (
		err error
		tx = storage.G_DB.Begin()
		uid = strconv.FormatInt(vo.AdminUser.Id, 10)
		buffer bytes.Buffer
		sql = "insert into `casbin_rule` (`p_type`,`v0`,`v1`,`v2`,`v5`) values"
	)
	// 用户表
	if isCreate {
		if err = tx.Create(&vo.AdminUser).Error; err != nil {
			goto ERR
		}
		uid = strconv.FormatInt(vo.AdminUser.Id, 10)
	} else {
		if err = tx.Model(&model.AdminUser{}).Save(&vo.AdminUser).Error; err != nil {
			goto ERR
		}
	}
	// 角色表
	// 需要删除的
	if vo.DeleteRids != nil && len(vo.DeleteRids) > 0 {
		if err = tx.Delete(&model.CasbinRule{}, "id in (?)", vo.DeleteRids).Error; err != nil {
			goto ERR
		}
	}
	// 需要添加的
	if vo.AddRoles != nil && len(vo.AddRoles) > 0 {
		if _, err := buffer.WriteString(sql); err != nil {
			return err
		}
		for i, v := range vo.AddRoles {
			if i == len(vo.AddRoles) - 1 {
				buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s');", "g", uid, v.V1, v.V2, v.V5))
			} else {
				buffer.WriteString(fmt.Sprintf("('%s','%s','%s','%s','%s'),", "g", uid, v.V1, v.V2, v.V5))
			}
		}
		if err = tx.Exec(buffer.String()).Error; err != nil {
			golog.Errorf("新建用户时，添加角色信息失败。错误：%s", err)
		}
	}
	tx.Commit()
	return nil
ERR:
	tx.Rollback()
	return err
}