package controller

import (
	"fmt"
	"message/dao/mysql"
	"message/dao/redis"
	"message/model"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

// 查詢全部留言
// get请求
func GetAll(c *gin.Context) {
	message, err := mysql.GetAllMessage()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// 查询留言，根据id
// get请求
func Get(c *gin.Context) {
	var message model.Message

	// 从redis中获取留言
	message, err := redis.GetMessage(redis.Client, c.Param("id"))
	if err != nil {
		fmt.Println("redis中没有此条留言, 从mysql中获取")
	}

	if err := mysql.GetMessage(&message, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "抱歉! 没有此条留言, 或者您输入的id不正确"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// 新增留言
// post请求
func Create(c *gin.Context) {
	var message model.Message

	if err := c.BindJSON(&message); err != nil { //c.BindJSON(&message) 从请求中获取数据并绑定到message结构体
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求数据格式不正确"})
		return
	}

	fmt.Println("Content", message.Content)

	if message.Content == "" || utf8.RuneCountInString(message.Content) >= 50 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "您输入的内容为空或者超过50个字符"})
		return
	}

	mysql.CreateMessage(&message)

	// 将留言存入redis
	redis.CreateMessage(redis.Client, fmt.Sprint(message.Id), message.User_Id, message.Content, message.CreatedAt.String())

	c.JSON(http.StatusCreated, gin.H{"message": message})
}

// 更新留言
// patch请求
func Update(c *gin.Context) {
	var message model.Message

	if err := c.BindJSON(&message); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求数据格式不正确"})
		return
	}

	fmt.Println("Content", message.Content)

	if message.Content == "" || utf8.RuneCountInString(message.Content) >= 50 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "您输入的内容为空或者超过50个字符"})
		return
	}

	if err := mysql.UpdateMessage(&message, c.PostForm("Content"), c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "抱歉! 没有此条留言, 或者您输入的id不正确"})
		return
	}

	// 删除redis中的留言
	redis.DeleteMessage(redis.Client, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// 删除留言
// delete请求
func Delete(c *gin.Context) {
	var message model.Message

	if err := mysql.DeleteMessage(&message, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "抱歉! 没有此条留言, 或者您输入的id不正确"})
		return
	}

	// 删除redis中的留言
	redis.DeleteMessage(redis.Client, c.Param("id"))

	c.JSON(http.StatusNoContent, gin.H{"message": "刪除留言成功"})
}
