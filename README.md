@Name:王继远  
@Date: 2024-05-20  
@Target:edgeness笔试

# 介绍
采用go+gin+mysql+redis实现的一个简单留言板，实现了留言的增删改查功能。
# 项目结构
```
.
├── conf
│   ├── mysql.yaml mysql配置文件
│   └── redis.yaml redis配置文件
├── controller
│   └── Controller.go 主要业务逻辑
├── dao
│   ├── mysql
│   │   ├── message.go 数据库操作
│   │   ├── message.sql 建立数据库表
│   │   └── mysql.go 数据库连接
│   └── redis
│       ├── message.go redis操作
│       └── redis.go redis连接
├── docker-compose.yaml docker-compose配置文件
├── Dockerfile
├── go.mod
├── go.sum
├── init.sql 初始化数据库
├── main.go 入口文件
├── model
│   └── Model.go 数据模型
├── README.md
└── router
    └── router.go 路由
```

# 项目运行
## 直接go build
由于项目比较简单，首先配置好mysql和redis（启两个docker容器），然后直接go build即可。
## docker-compose
