
--- Menu: 设备管理/角色管理
SELECT menu_id,path into @sysrole_parent_menu_id,@sysrole_parent_menu_path FROM sys_menu where parent_id=0 and menu_name='设备管理';
insert into sys_menu (menu_name, parent_id, order_num, path, component, is_frame, is_cache, menu_type, visible, status, perms, icon, create_by, create_time, update_by, update_time, remark)
 values('角色管理', @sysrole_parent_menu_id, '1', 'sysrole', 'sysrole/index', 1, 0, 'C', '0', '0', '', 'guide', 'admin', sysdate(), '', null, '角色管理');

--- Menu: 实时监控/动力系统
SELECT menu_id,path into @testcat_parent_menu_id,@testcat_parent_menu_path FROM sys_menu where parent_id=0 and menu_name='实时监控';
insert into sys_menu (menu_name, parent_id, order_num, path, component, is_frame, is_cache, menu_type, visible, status, perms, icon, create_by, create_time, update_by, update_time, remark)
 values('动力系统', @testcat_parent_menu_id, '1', 'testcat', 'testcat/index', 1, 0, 'C', '0', '0', '', 'guide', 'admin', sysdate(), '', null, '动力系统');
