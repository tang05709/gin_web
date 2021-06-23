-- 插入微信用户菜单
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '微信用户管理', 0, '/admin/module/member', 0, 1, 'GET');

SET @pid = LAST_INSERT_ID();

INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '微信用户列表', @pid, '/admin/module/members', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑微信用户', @pid, '/admin/module/member/edit', 0, 1, 'GET');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '编辑微信用户API', @pid, '/admin/module/member/save', 1, 0, 'POST');
INSERT INTO `s_menus`(`created_at`, `name`, `parent_id`, `router`, `type`, `visible`, `method`) VALUES (now(), '删除微信用户API', @pid, '/admin/module/member/del', 1, 0, 'POST');