package main

import (
	"fmt"
	"message/dao/mysql"
	"message/dao/redis"
	"message/router"
)

func main() {
	//初始化mysql
	if err := mysql.InitMySql(); err != nil {
		panic(err)
	}

	//建立表
	mysql.Connect.Table("message")

	//初始化redis
	if err := redis.InitRedis(); err != nil {
		panic(err)
	}

	//注册路由
	r := router.SetRouter()

	//启动服务
	fmt.Println("请到127.0.0.0.1:8081...访问服务")
	r.Run("127.0.0.1:8081")
}
