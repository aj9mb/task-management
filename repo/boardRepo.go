package repo

import (
	"log"

	"github.com/aj9mb/task-management/dbmg"
	"github.com/aj9mb/task-management/model"
	_ "github.com/go-sql-driver/mysql"
)

func BoardAdd(b *model.Board) (*model.Board, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare("INSERT INTO board(name, created_by) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	res, err := stmt.Exec(b.Name, b.CreatedBy)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	b.Id = id
	return b, err
}

func BoardUserAdd(b *model.BoardUser) (int64, error) {
	db := dbmg.GetDb()
	stmt, err := db.Prepare("INSERT INTO board_user (board_id, user_id, added_by, active) VALUES(?, ?, ?, 1) ON DUPLICATE KEY UPDATE added_by = VALUES(added_by), active = VALUES(active)")
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	res, err := stmt.Exec(b.BoardId, b.UserId, b.AddedBy)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, err
}

func BoardListGet(userId int64) (*[]model.Board, error) {
	db := dbmg.GetDb()
	boardlist := make([]model.Board, 0)
	stmt, err := db.Prepare("SELECT b.board_id, b.name from board b JOIN board_user bu on b.board_id = bu.board_id where bu.user_id = ?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := stmt.Query(userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		board := new(model.Board)
		err := rows.Scan(&board.Id, &board.Name)
		if err != nil {
			log.Fatal(err)
		}
		boardlist = append(boardlist, *board)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &boardlist, err
}
