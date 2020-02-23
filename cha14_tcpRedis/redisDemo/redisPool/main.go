package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局的pool,
var pool *redis.Pool

//当启动程序是就初始化连接池pool
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   //和数据库最大连接数,0表示无限制
		IdleTimeout: 300, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接匿名函数,用于连接redis server,返回连接
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//pool.Close() //测试关闭连接池
	//从pool内取出一个连接
	conn := pool.Get()
	defer conn.Close() //用完需要关闭以便放回至连接池

	strs, err := redis.Strings(conn.Do("HMGet", "user1", "name", "age", "addr", "score"))
	if err != nil {
		fmt.Println("conn.Do err:", err)
	}
	for _, v := range strs {
		fmt.Println(v)
	}
}
