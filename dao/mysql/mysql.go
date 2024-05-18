package mysql

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connect *gorm.DB

type conf struct {
	Host     string `yaml:"host"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func (c *conf) getConf() *conf {
	//获取yaml配置文件
	yamlFile, err := os.ReadFile("dao/mysql/mysql.yaml")

	//如果读取文件失败，打印错误信息
	if err != nil {
		fmt.Println(err.Error())
	}

	//将yaml配置文件解析到结构体中
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

// 初始化mysql
func InitMySql() (err error) {
	var c conf

	//获取yaml配置文件
	conf := c.getConf()

	//将配置文件中的信息拼接成dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DbName,
	)

	//连接数据库
	Connect, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	return
}
