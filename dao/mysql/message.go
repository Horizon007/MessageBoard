package mysql

import (
	"message/model"
)

// 查询全部留言
func GetAllMessage() (message []*model.Message, err error) {
	err = Connect.Find(&message).Error
	return
}

// 根据 id 查询留言
func GetMessage(message *model.Message, id string) (err error) {
	err = Connect.Where("id=?", id).First(&message).Error
	return
}

// 新增留言
func CreateMessage(message *model.Message) (err error) {
	err = Connect.Create(&message).Error
	return
}

// 根据 id 更新留言
func UpdateMessage(message *model.Message, content, id string) (err error) {
	err = Connect.Where("id=?", id).First(&message).Update("content", content).Error
	return
}

// 根据 id 删除留言
func DeleteMessage(message *model.Message, id string) (err error) {
	err = Connect.Where("id=?", id).First(&message).Delete(&message).Error
	return
}
