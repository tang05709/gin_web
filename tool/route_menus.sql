-- 插入线路菜单
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '线路管理', 0, '/admin/module/route', 0, 1, 'GET');

SET @pid = LAST_INSERT_ID();

INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '线路列表', @pid, '/admin/module/routes', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑线路', @pid, '/admin/module/route/edit', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑线路API', @pid, '/admin/module/route/save', 1, 0, 'POST');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '删除线路API', @pid, '/admin/module/route/del', 1, 0, 'POST');