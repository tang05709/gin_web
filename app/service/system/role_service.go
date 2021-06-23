package system

import (
	"festival/app/common/cache"
	"festival/app/common/db"
	"festival/app/common/utils/page"
	"festival/app/model/system"

	"gorm.io/gorm"
)

// 获取角色列表
func RoleList(req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&system.SysRole{})
	list := make([]system.SysRole, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Preload("Menus").Order("created_at desc").Find(&list).Error
	return err, list, total
}

func GetRole(id uint) (u system.SysRole, err error) {
	var item system.SysRole
	item.Id = uint(id)
	err = db.Global().Preload("Menus").First(&item).Error
	if err != nil {
		return item, err
	}
	return item, err
}

func DeleteRole(item system.SysRole) (err error) {
	err = db.Global().Delete(&item).Error
	cache.Instance().Flush()
	return err
}

func RoleAll() (err error, rows []system.SysRole) {
	query := db.Global().Model(&system.SysRole{})
	rows = make([]system.SysRole, 0)
	err = query.Order("created_at desc").Find(&rows).Error
	return err, rows
}

// 增加用户角色
func AddRole(item system.SysRole) (err error) {
	err = db.Global().Create(&item).Error
	return err
}

// 更新用户角色
func UpdateRole(item system.SysRole) (err error) {
	err = db.Global().Transaction(func(tx *gorm.DB) error {
		// omit 排除更新内容
		txErr := tx.Omit("id", "Menus").Session(&gorm.Session{FullSaveAssociations: true}).Updates(&item).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Model(&item).Association("Menus").Replace(&item.Menus)
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}
