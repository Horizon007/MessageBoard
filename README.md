> @Name:王继远 
> @Date: 2024-05-20
> @Target:edgeness笔试

# 介绍
采用go+gin+mysql+redis实现的一个简单留言板，实现了留言的增删改查功能。
# 项目结构
```
.
├── controller
│   └── Controller.go //增删改查功能的实现
├── dao
│   ├── mysql
│   │   ├── message.go   //mysql操作
│   │   ├── message.sql  //mysql建表语句
│   │   ├── mysql.go     //mysql初始化连接
│   │   └── mysql.yaml   //mysql配置文件
│   └── redis
│       ├── message.go   //redis操作
│       ├── redis.go     //redis初始化连接
│       └── redis.yaml   //redis配置文件
├── go.mod
├── go.sum
├── main.go              //程序入口
├── model
│   └── Model.go         //数据模型
└── router
    └── router.go        //路由
```

# 项目运行
## 直接go build
由于项目比较简单，首先配置好mysql和redis（启两个docker容器），然后直接go build即可。
## docker-compose
