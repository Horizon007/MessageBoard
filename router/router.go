package router

import (
	"message/controller"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {

	//发布模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1 := r.Group("api/v1") //路由组
	{
		//新增留言
		v1.POST("/message", controller.Create)
		//查询全部留言
		v1.GET("/message", controller.GetAll)
		//根据 id 查询留言
		v1.GET("/message/:id", controller.Get)
		//根据 id 更新留言
		v1.PATCH("/message/:id", controller.Update)
		//根据 id 删除留言
		v1.DELETE("/message/:id", controller.Delete)
	}
	return r
}
