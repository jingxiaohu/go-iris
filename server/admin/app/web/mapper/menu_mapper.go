package mapper

import (
	"database/sql"
	"goiris/admin/app/web/vo"
	"goiris/common/model"
	"goiris/common/storage"
)

type MenuMapper struct {
}

func (mm *MenuMapper) Insert(menu *model.Menu) error {
	return storage.G_DB.Create(&menu).Error
}

func (mm *MenuMapper) UpdateOne(menu *model.Menu) error {
	return storage.G_DB.Model(&model.Menu{}).Save(&menu).Error
}

func (mm *MenuMapper) Delete(mids, mMids []int64) error {
	var (
		err error
		tx = storage.G_DB.Begin()
	)
	if err = tx.Delete(model.Menu{}, "id in (?)", mids).Error; err != nil {
		goto ERR
	}
	if err = tx.Delete(model.MenuMeta{}, "id in (?)", mMids).Error; err != nil {
		goto ERR
	}
	tx.Commit()
	return nil
ERR:
	tx.Rollback()
	return err
}

func (mm *MenuMapper) AllMenu() ([]*model.Menu, error) {
	var (
		err    error
		result = make([]*model.Menu, 0)
	)
	if err = storage.G_DB.Preload("Children.Meta").
		Preload("Meta").
		Order("sort").
		Find(&result, "is_sub = 0").Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (mm *MenuMapper) OwnMenu(uid string) ([]int64, error) {
	var (
		err     error
		rows    *sql.Rows
		menuId  int64
		menuIds []int64
	)
//	SELECT mid FROM role_menu
//	WHERE rid in (
//		SELECT id FROM casbin_rule
//	WHERE p_type='p' AND v0 in (
//		SELECT v1 FROM casbin_rule WHERE p_type='g' AND v0=?
//))
	rows, err = storage.G_DB.Raw(`
SELECT mid FROM role_menu
WHERE role_key in (
SELECT v1 FROM casbin_rule WHERE p_type='g' AND v0=? GROUP BY v1)
`, uid).Rows();
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&menuId)
		menuIds = append(menuIds, menuId)
	}
	return menuIds, nil
}

func (mm *MenuMapper) FindList(vo *vo.MenuVO) (int64, []*model.Menu, error) {
	var (
		err    error
		total  int64
		result = make([]*model.Menu, 0)
	)
	if err = storage.G_DB.Model(&model.Menu{}).Where("is_sub = 0").Count(&total).Error; err != nil {
		goto ERR
	}
	if err = storage.G_DB.
		Preload("Children.Meta").
		Preload("Meta").
		Order("sort").
		Offset(vo.Start).Limit(vo.Size).
		Find(&result, "is_sub = 0").Error; err != nil {
		goto ERR
	}
	return total, result, nil
ERR:
	return 0, nil, err
}

func (mm *MenuMapper) ParentMenus() ([]*model.Menu, error) {
	var (
		err    error
		result = make([]*model.Menu, 0)
	)
	if err = storage.G_DB.
		Preload("Meta").
		Order("sort").
		Find(&result, "is_sub = 0").Error; err != nil {
		goto ERR
	}
	return result, nil
ERR:
	return nil, err
}