package models

import (
	"fmt"
	u "github.com/MrTBit/goForumEngine/utils"
	"github.com/jinzhu/gorm"
)

type Thread struct {
	gorm.Model
	Threadname        string `json:"threadname"`
	ThreadID          uint   `json:"threadID"`
	ThreadDescription string `json:"threaddescription"`
	BoardID           uint   `json:"boardID"`
}

//validate new thread details
func (thread *Thread) Validate() (map[string]interface{}, bool) {
	if len(thread.Threadname) < 2 {
		return u.Message(false, "Thread name is required"), false
	}

	if len(thread.ThreadDescription) < 2 {
		return u.Message(false, "Thread description required"), false
	}

	if thread.BoardID <= 0 {
		return u.Message(false, "boardID required"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (thread *Thread) Create() map[string]interface{} {
	if resp, ok := thread.Validate(); !ok {
		return resp
	}

	temp := &Thread{}
	err := GetDB().Table("threads").Order("threadID desc").First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection Error. Please retry")
	}

	if temp.ThreadID <= 0 {
		thread.ThreadID = 1
	} else {
		thread.ThreadID = temp.ThreadID + 1
	}

	GetDB().Create(thread)

	if thread.ID <= 0 {
		return u.Message(false, "Failed to create thread, connection error.")
	}

	response := u.Message(true, "Thread created.")
	response["thread"] = thread
	return response
}

func GetThreads(boardid uint) []*Thread {
	threads := make([]*Thread, 0)
	err := GetDB().Table("threads").Where("boardID = ?", boardid).Find(&threads).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return threads
}

func GetThread(threadid uint) *Thread {
	thread := &Thread{}
	err := GetDB().Table("threads").Where("threadID = ?", threadid).First(thread).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return thread
}
