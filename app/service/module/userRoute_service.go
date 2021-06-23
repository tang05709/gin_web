package module

import (
	"festival/app/common/db"
	"festival/app/common/utils/page"
	"festival/app/model/module"
	"gorm.io/gorm"
)

// 点亮线路表
// power by 7be.cn

// 获取菜单列表
func ModUserRouteList(req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&module.ModUserRoute{}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, account, user_name, gender, avatar")
	}).Preload("Route")
	list := make([]module.ModUserRoute, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Order("created_at desc").Find(&list).Error
	return err, list, total
}

func GetModUserRoute(id uint) (u module.ModUserRoute, err error) {
	var item module.ModUserRoute
	item.Id = id
	err = db.Global().First(&item).Error
	if err != nil {
		return item, err
	}
	return item, err
}

func DeleteModUserRoute(item module.ModUserRoute) (err error) {
	err = db.Global().Delete(&item).Error
	return err
}

// 增加用户角色
func AddModUserRoute(item *module.ModUserRoute) (err error) {
	err = db.Global().Create(&item).Error
	return err
}

// 更新用户角色
func UpdateModUserRoute(item *module.ModUserRoute) (err error) {
	err = db.Global().Select("*").Omit("id", "created_at").Updates(&item).Error
	return err
}

//根据用户获取路线
func UserRouteListByUser(userId uint, req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&module.ModUserRoute{}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, account, user_name, gender, avatar")
	}).Preload("Route").Where("userId = ?", userId)
	list := make([]module.ModUserRoute, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Order("created_at desc").Find(&list).Error
	return err, list, total
}