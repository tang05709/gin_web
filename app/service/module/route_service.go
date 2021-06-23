package module

import (
	"festival/app/common/db"
	"festival/app/common/utils/page"
	"festival/app/model/module"
)

// 线路
// power by 7be.cn

// 获取菜单列表
func ModRouteList(req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&module.ModRoute{})
	list := make([]module.ModRoute, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Order("created_at desc").Find(&list).Error
	return err, list, total
}

func GetModRoute(id uint) (u module.ModRoute, err error) {
	var item module.ModRoute
	item.Id = id
	err = db.Global().First(&item).Error
	if err != nil {
		return item, err
	}
	return item, err
}

func DeleteModRoute(item module.ModRoute) (err error) {
	err = db.Global().Delete(&item).Error
	return err
}

// 增加用户角色
func AddModRoute(item *module.ModRoute) (err error) {
	err = db.Global().Create(&item).Error
	return err
}

// 更新用户角色
func UpdateModRoute(item *module.ModRoute) (err error) {
	err = db.Global().Select("*").Omit("id", "created_at").Updates(&item).Error
	return err
}