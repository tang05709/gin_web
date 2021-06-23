-- 插入点亮线路表菜单
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '点亮线路表管理', 0, '/admin/module/userRoute', 0, 1, 'GET');

SET @pid = LAST_INSERT_ID();

INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '点亮线路表列表', @pid, '/admin/module/userRoutes', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑点亮线路表', @pid, '/admin/module/userRoute/edit', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑点亮线路表API', @pid, '/admin/module/userRoute/save', 1, 0, 'POST');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '删除点亮线路表API', @pid, '/admin/module/userRoute/del', 1, 0, 'POST');