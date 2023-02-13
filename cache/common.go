package cache

import (
	"Minimalist_TikTok/model/websocket"
	"Minimalist_TikTok/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"strconv"
	"time"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("redis配置文件读取错误，请检查文件路径")
	}
	LoadRedis(file)
	Redis()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) //string 2 uint64
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		util.LogrusObj.Info(err)
		panic(err)
	}
	RedisClient = client
}

func RedisInsert(database, id string, content string, read uint, expire int64) {
	ctx := context.Background()
	con := websocket.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	marshal, _ := json.Marshal(con)
	//fmt.Println(marshal)
	fmt.Println("id=", id)
	//有个稀里糊涂的bug
	RedisClient.LPush(ctx, id, marshal)

}
