package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 密码，如果没有则留空
		DB:       0,                // 使用的数据库编号
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("无法连接到 Redis:", err)
		return
	}
	fmt.Println("连接成功:", pong)

	// 创建索引
	_, err = rdb.Do(ctx, "FT.CREATE", "myIndex", "ON", "HASH", "PREFIX", "1", "doc:",
		"SCHEMA", "title", "TEXT", "WEIGHT", "5.0", "body", "TEXT").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("索引创建成功")

	// 添加文档
	_, err = rdb.HSet(ctx, "doc:1", "title", "Redis Tutorial", "body", "Learn how to use Redis with Go").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("文档添加成功")

	// 搜索文档
	searchResult, err := rdb.Do(ctx, "FT.SEARCH", "myIndex", "Redis").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("搜索结果:", searchResult)
}
