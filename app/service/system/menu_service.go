package system

import (
	"festival/app/common/cache"
	"festival/app/common/db"
	"festival/app/common/utils/gconv"
	"festival/app/common/utils/page"
	"festival/app/model"
	"festival/app/model/system"
	systemModel "festival/app/model/system"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取菜单列表
func MenuList(req *page.Paging) (err error, rows interface{}, total int64) {
	query := db.Global().Model(&system.SysMenu{})
	list := make([]system.SysMenu, 0)
	err = query.Count(&total).Error
	err = query.Limit(req.Pagesize).Offset(req.StartNum).Order("created_at desc").Find(&list).Error
	return err, list, total
}

// 获取菜单树列表
func MenuTreeList() (err error, rows []system.SysMenu) {
	query := db.Global().Model(&system.SysMenu{})
	rows = make([]system.SysMenu, 0)
	err = query.Order("created_at desc").Find(&rows).Error
	menus, err := getMenuTree(rows)
	return err, menus
}

func MenuAll() (err error, rows []system.SysMenu) {
	query := db.Global().Model(&system.SysMenu{})
	rows = make([]system.SysMenu, 0)
	err = query.Order("created_at desc").Find(&rows).Error
	return err, rows
}

func MenuInIds(ids []string) (err error, rows []system.SysMenu) {
	query := db.Global().Model(&system.SysMenu{})
	rows = make([]system.SysMenu, 0)
	err = query.Where("id in ?", ids).Find(&rows).Error
	return err, rows
}

func GetMenu(id uint) (u system.SysMenu, err error) {
	var item system.SysMenu
	item.Id = uint(id)
	err = db.Global().First(&item).Error
	if err != nil {
		return item, err
	}
	return item, err
}

func DeleteMenu(item system.SysMenu) (err error) {
	err = db.Global().Delete(&item).Error
	ClearAllMenus()
	return err
}

// 增加用户角色
func AddMenu(item *system.SysMenu) (err error) {
	err = db.Global().Create(&item).Error
	return err
}

// 更新用户角色
func UpdateMenu(item *system.SysMenu) (err error) {
	err = db.Global().Select("*").Omit("id", "created_at").Updates(&item).Error
	return err
}

func ClearRoleMenus(roleId uint) {
	key := model.AUTH_CACHE + gconv.String(roleId)
	cache.Instance().Delete(key)
}

func ClearAllMenus() {
	cache.Instance().Flush()
}

// 查询用户的菜单
func GetUserMenus(c *gin.Context) ([]systemModel.SysMenu, error) {
	user := GetProfile(c)
	if user == nil {
		return nil, errors.New("获取失败")
	}
	return getRoleApis(user.RoleId)
}

func GetUserAuth(c *gin.Context) (map[string]systemModel.SysMenu, error) {
	user := GetProfile(c)
	if user == nil {
		return nil, errors.New("获取失败")
	}
	return getRoleAuth(user.RoleId)
}

// 查询角色的权限
func getRoleAuth(roleId uint) (map[string]systemModel.SysMenu, error) {
	//从缓存读取
	key := model.AUTH_CACHE + gconv.String(roleId)
	tmp, have := cache.Instance().Get(key)
	if have && tmp != nil {
		rs, ok := tmp.(map[string]systemModel.SysMenu)
		if ok {
			return rs, nil
		}
	}
	var role systemModel.SysRole
	var menus []systemModel.SysMenu
	role.Id = roleId
	// 从数据库中读取
	err := db.Global().Model(role).Preload("Menus").First(&role).Error
	if err != nil {
		return nil, err
	}
	menus = role.Menus
	// 插入所有的 router - menu
	routerMap := make(map[string]systemModel.SysMenu, 0)
	for _, v := range menus {
		k := v.Method + "-" + v.Router
		routerMap[k] = v
	}
	if routerMap != nil {
		cache.Instance().Set(key, routerMap, 7*24*time.Hour)
	}
	return routerMap, err
}

// 查询角色的菜单
func getRoleApis(roleId uint) ([]systemModel.SysMenu, error) {
	//从缓存读取
	key := model.MENU_CACHE + gconv.String(roleId)
	tmp, have := cache.Instance().Get(key)

	if have && tmp != nil {
		rs, ok := tmp.([]systemModel.SysMenu)
		if ok {
			return rs, nil
		}
	}
	var role systemModel.SysRole
	var menus []systemModel.SysMenu
	role.Id = roleId
	// 从数据库中读取
	err := db.Global().Model(role).Preload("Menus").First(&role).Error
	if err != nil {
		return nil, err
	}
	menus = role.Menus

	menus, err = getMenuTree(menus)

	if err == nil {
		cache.Instance().Set(key, menus, 7*24*time.Hour)
	}
	return menus, err
}

func getMenuTree(menus []systemModel.SysMenu) (m []systemModel.SysMenu, err error) {
	// 对菜单进行循环
	treeMap := make(map[int][]systemModel.SysMenu)
	for _, v := range menus {
		// if v.Type == 0 {
		treeMap[int(v.ParentId)] = append(treeMap[int(v.ParentId)], v)
		// }
	}
	// 获取根节点
	menus = treeMap[0]
	// 迭代获取每个根节点下的子节点
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

func getChildrenList(menu *systemModel.SysMenu, treeMap map[int][]systemModel.SysMenu) (err error) {
	menu.Children = treeMap[int(menu.Id)]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
