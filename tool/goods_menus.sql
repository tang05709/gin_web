-- 插入礼品表菜单
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '礼品表管理', 0, '/admin/module/goods', 0, 1, 'GET');

SET @pid = LAST_INSERT_ID();

INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '礼品表列表', @pid, '/admin/module/goodss', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑礼品表', @pid, '/admin/module/goods/edit', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑礼品表API', @pid, '/admin/module/goods/save', 1, 0, 'POST');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '删除礼品表API', @pid, '/admin/module/goods/del', 1, 0, 'POST');