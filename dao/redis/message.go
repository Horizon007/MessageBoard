package redis

import (
	"context"
	"message/model"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

const layout = "2006-01-02 15:04:05.99 -0700 MST"

func CreateMessage(rdb *redis.Client, id string, User_Id int, content string, c_ts string) error {
	err := rdb.HSet(ctx, "message:"+id, "user_id", User_Id, "content", content, "c_ts", c_ts).Err()
	return err
}

func GetMessage(rdb *redis.Client, id string) (model.Message, error) {
	val, err := rdb.HGetAll(ctx, "message:"+id).Result()

	//将这个map转换为Message结构体
	var message model.Message
	message.Id, _ = strconv.Atoi(id)
	message.User_Id, _ = strconv.Atoi(val["user_id"])
	message.Content = val["content"]
	message.CreatedAt, _ = time.Parse(layout, val["timestamp"])

	return message, err
}

func DeleteMessage(rdb *redis.Client, id string) error {
	err := rdb.Del(ctx, "message:"+id).Err()
	return err
}
