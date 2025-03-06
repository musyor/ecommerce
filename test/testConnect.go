package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
)

func main() {

}

// 测试MySQL连接
func testMySQLConnect() error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommerce")
	if err != nil {
		return fmt.Errorf("数据库连接失败：%v", err)
	}
	defer func(db *sql.DB) error {
		err := db.Close()
		if err != nil {
			return fmt.Errorf("数据库连接释放失败：%v,err")
		}
		return nil
	}(db)

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("数据库相应失败：%v", err)
	}

	fmt.Println("数据库连接成功...")
	return nil
}

func testRedisConnect() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:3306",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("redis ping失败：%v", err)
	}
	fmt.Println("redis 连接成功：", pong)
	return nil
}

func testKafkaConnect() error {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		return fmt.Errorf("kafka连接失败：%v", err)
	}
	defer conn.Close()
	fmt.Println("成功连接Kafak")
	return nil
}
