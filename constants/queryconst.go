package constants

var (
	BOARD_ADD       = "INSERT INTO board(name, created_by) VALUES(?, ?)"
	BOARD_USER_ADD  = "INSERT INTO board_user (board_id, user_id, added_by, active) VALUES(?, ?, ?, 1) ON DUPLICATE KEY UPDATE added_by = VALUES(added_by), active = VALUES(active)"
	BOARD_LIST_GET  = "SELECT b.board_id, b.name from board b JOIN board_user bu on b.board_id = bu.board_id and bu.active = 1 where bu.user_id = ?"
	USER_ADD        = "INSERT INTO `user`(user_name, password, name, active) VALUES(?,?,?,1)"
	USER_NAME_CHECK = "SELECT user_name FROM `user` WHERE user_name in ("
	USER_GET        = "SELECT user_id, user_name, password, name FROM `user` WHERE active = 1 and user_name = ?"
	TASK_ADD        = "INSERT INTO task (board_id, added_by, assigned_to, task_desc, active) VALUES(?,?,?,?,1)"
)
