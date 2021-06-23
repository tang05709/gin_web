-- 插入兑换礼品表菜单
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '兑换礼品表管理', 0, '/admin/module/userGoods', 0, 1, 'GET');

SET @pid = LAST_INSERT_ID();

INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '兑换礼品表列表', @pid, '/admin/module/userGoodss', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑兑换礼品表', @pid, '/admin/module/userGoods/edit', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑兑换礼品表API', @pid, '/admin/module/userGoods/save', 1, 0, 'POST');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '删除兑换礼品表API', @pid, '/admin/module/userGoods/del', 1, 0, 'POST');