package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
)

var Client *redis.Client

type conf struct {
	Host      string `yaml:"host"`
	Db        int    `yaml:"db"`
	Pool_size int    `yaml:"pool_size"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := os.ReadFile("./conf/redis.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func InitRedis() (err error) {
	var c conf
	conf := c.getConf()

	Client = redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		DB:       conf.Db,
		PoolSize: conf.Pool_size,
	})

	_, err = Client.Ping(context.Background()).Result()
	return
}
