package models

import (
	"fmt"
	u "github.com/MrTBit/goForumEngine/utils"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Commentbody   string `json:"commentbody"`
	CommentID     uint   `json:"commentID"`
	CommentAuthor string `json:"commentauthor"`
	ThreadID      uint   `json:"threadID"`
}

//validate new thread details
func (comment *Comment) Validate() (map[string]interface{}, bool) {
	if len(comment.CommentAuthor) < 2 {
		return u.Message(false, "Comment author is required"), false
	}

	if len(comment.Commentbody) < 2 {
		return u.Message(false, "Comment body required"), false
	}

	if comment.ThreadID <= 0 {
		return u.Message(false, "threadID required"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (comment *Comment) Create() map[string]interface{} {
	if resp, ok := comment.Validate(); !ok {
		return resp
	}

	temp := &Comment{}
	err := GetDB().Table("comments").Order("comment_id desc").First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection Error. Please retry")
	}

	if temp.CommentID <= 0 {
		comment.CommentID = 1
	} else {
		comment.CommentID = temp.CommentID + 1
	}

	GetDB().Create(comment)

	if comment.ID <= 0 {
		return u.Message(false, "Failed to create thread, connection error.")
	}

	response := u.Message(true, "Comment created.")
	response["comment"] = comment
	return response
}

func GetComments(threadid uint) []*Comment {
	comments := make([]*Comment, 0)
	err := GetDB().Table("comments").Where("thread_id = ?", threadid).Find(&comments).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return comments
}

func GetComment(commentid uint) *Comment {
	comment := &Comment{}
	err := GetDB().Table("comments").Where("comment_id = ?", commentid).First(comment).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return comment
}
