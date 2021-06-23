package module

import (
	"festival/app/common/db"
	"festival/app/common/utils/page"
	"festival/app/model/module"
	"gorm.io/gorm"
)

// 兑换礼品表
// power by 7be.cn

// 获取菜单列表
func ModUserGoodsList(req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&module.ModUserGoods{}).Preload("User").Preload("Goods")
	list := make([]module.ModUserGoods, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Order("created_at desc").Find(&list).Error
	return err, list, total
}

func GetModUserGoods(id uint) (u module.ModUserGoods, err error) {
	var item module.ModUserGoods
	item.Id = id
	err = db.Global().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, account, user_name, gender, avatar")
	}).Preload("Goods").First(&item).Error
	if err != nil {
		return item, err
	}
	return item, err
}

func DeleteModUserGoods(item module.ModUserGoods) (err error) {
	err = db.Global().Delete(&item).Error
	return err
}

// 增加用户角色
func AddModUserGoods(item *module.ModUserGoods) (err error) {
	err = db.Global().Create(&item).Error
	return err
}

// 更新用户角色
func UpdateModUserGoods(item *module.ModUserGoods) (err error) {
	err = db.Global().Select("*").Omit("id", "created_at").Updates(&item).Error
	return err
}

//根据用户获取路线
func UserGoodsListByUser(userId uint, req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&module.ModUserGoods{}).Where("userId = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, account, user_name, gender, avatar")
	}).Preload("Goods")
	list := make([]module.ModUserGoods, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Order("created_at desc").Find(&list).Error
	return err, list, total
}

func UpdateModUserGoodsStatus(id uint, status int) (err error) {
	err = db.Global().Model(&module.ModUserGoods{}).Where("id = ?", id).Update("status", status).Error
	return err
}