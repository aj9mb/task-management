package constants

var (
	BOARD_ADD         = "INSERT INTO board(name, created_by) VALUES(?, ?)"
	BOARD_USER_ADD    = "INSERT INTO board_user (board_id, user_id, added_by, active) VALUES(?, ?, ?, 1) ON DUPLICATE KEY UPDATE added_by = VALUES(added_by), active = VALUES(active)"
	BOARD_LIST_GET    = "SELECT b.board_id, b.name from board b JOIN board_user bu on b.board_id = bu.board_id and bu.active = 1 where bu.user_id = ?"
	USER_ADD          = "INSERT INTO `user`(user_name, password, name, active) VALUES(?,?,?,1)"
	USER_NAME_CHECK   = "SELECT user_name FROM `user` WHERE user_name in ("
	USER_GET          = "SELECT user_id, user_name, password, name FROM `user` WHERE active = 1 and user_name = ?"
	TASK_ADD          = "INSERT INTO task (board_id, added_by, assigned_to, task_desc, status, active) VALUES(?,?,?,?,0,1)"
	CHECK_BOARD_EXIST = "SELECT 1 from board where board_id = ?"
	GET_TASK_LIST     = "SELECT t.task_id, t.board_id, t.added_by, u1.name as added_by_name, t.assigned_to, u2.name as assigned_to_name, t.task_desc, t.status, t.create_dt, t.last_update  from task t join board b on b.board_id = t.board_id join board_user bu on bu.board_id = b.board_id  join user u1 on u1.user_id = t.added_by join user u2 on u2.user_id = t.assigned_to  where t.board_id = ? and t.active = 1 and bu.user_id = ?"
)
