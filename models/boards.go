package models

import (
	"fmt"
	u "github.com/MrTBit/goForumEngine/utils"
	"github.com/jinzhu/gorm"
)

type Board struct {
	gorm.Model
	Boardname        string `json:"boardname"`
	BoardID          uint   `json:"boardID"`
	BoardDescription string `json:"boarddescription"`
}

//validate new board details
func (board *Board) Validate() (map[string]interface{}, bool) {
	if len(board.Boardname) < 2 {
		return u.Message(false, "Board name is required"), false
	}

	//Board name must be unique
	temp := &Board{}

	//check for errors and duplicate board names
	err := GetDB().Table("boards").Where("boardname = ?", board.Boardname).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Boardname != "" {
		return u.Message(false, "Board name already taken."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (board *Board) Create() map[string]interface{} {
	if resp, ok := board.Validate(); !ok {
		return resp
	}

	temp := &Board{}
	err := GetDB().Table("boards").Order("board_id desc").First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection Error. Please retry")
	}

	if temp.BoardID <= 0 {
		board.BoardID = 1
	} else {
		board.BoardID = temp.BoardID + 1
	}

	GetDB().Create(board)

	if board.ID <= 0 {
		return u.Message(false, "Failed to create board, connection error.")
	}

	response := u.Message(true, "Board created.")
	response["board"] = board
	return response
}

func GetBoards() []*Board {
	boards := make([]*Board, 0)
	err := GetDB().Table("boards").Find(&boards).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return boards
}
