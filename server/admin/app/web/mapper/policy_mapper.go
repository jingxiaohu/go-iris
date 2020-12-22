package mapper

import (
	"goiris/admin/app/web/vo"
	"goiris/common/model"
	"goiris/common/storage"
)

type PolicyMapper struct {
}

func (pm *PolicyMapper) Delete(ids []int64) error {
	return storage.G_DB.Delete(model.CasbinRule{}, "id in (?)", ids).Error
}

func (pm *PolicyMapper) UpdateOne(role *model.CasbinRule) error {
	return storage.G_DB.Model(&model.CasbinRule{}).Update(&role).Error
}

func (pm *PolicyMapper) Table(vo *vo.RoleVO) (int64, []*model.CasbinRule, error) {
	var (
		err    error
		total  int64
		result = make([]*model.CasbinRule, 0)
		db = storage.G_DB.Offset(vo.Start).Limit(vo.Size)
	)
	if err = storage.G_DB.Model(&model.CasbinRule{}).Where("p_type='p'").Count(&total).Error; err != nil {
		goto ERR
	}
	if vo.Domain == "" || vo.Domain == "all" {
		if err = db.Find(&result, "p_type='p'").Error; err != nil {
			goto ERR
		}
	} else {
		if err = db.Find(&result, "p_type='p' AND v1=?", vo.Domain).Error; err != nil {
			goto ERR
		}
	}
	return total, result, nil
ERR:
	return 0, nil, err
}

func (pm *PolicyMapper) All() ([]*model.CasbinRule, error) {
	var (
		err    error
		result = make([]*model.CasbinRule, 0)
	)
	if err = storage.G_DB.Find(&result, "p_type='p'").Error; err != nil {
		goto ERR
	}
	return result, nil
ERR:
	return nil, err
}