package module

import (
	"festival/app/common/db"
	"festival/app/common/utils/page"
	"festival/app/model/module"
)

// 礼品表
// power by 7be.cn

// 获取菜单列表
func ModGoodsList(req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&module.ModGoods{})
	list := make([]module.ModGoods, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Order("created_at desc").Find(&list).Error
	return err, list, total
}

func GetModGoods(id uint) (u module.ModGoods, err error) {
	var item module.ModGoods
	item.Id = id
	err = db.Global().First(&item).Error
	if err != nil {
		return item, err
	}
	return item, err
}

func DeleteModGoods(item module.ModGoods) (err error) {
	err = db.Global().Delete(&item).Error
	return err
}

// 增加用户角色
func AddModGoods(item *module.ModGoods) (err error) {
	err = db.Global().Create(&item).Error
	return err
}

// 更新用户角色
func UpdateModGoods(item *module.ModGoods) (err error) {
	err = db.Global().Select("*").Omit("id", "created_at").Updates(&item).Error
	return err
}