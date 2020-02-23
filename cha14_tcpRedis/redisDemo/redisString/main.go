package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//引入Redis包

//通过Go向Redis写入及读取数据
func main() {
	//1. 连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if err != nil {
		fmt.Printf("redis.Dial err : %v\n", err)
		return
	}
	fmt.Println("connect OK", conn)

	//2. 通过Go向redis写入数据 string
	_, err = conn.Do("Set", "name", "Tom猫")
	if err != nil {
		fmt.Printf("conn.Do Set err : %v\n", err)
		return
	}
	fmt.Println("Set OK")
	//3. 通过Go向redis读取数据 string
	//返回的是interface{}空接口,需要转换成对应格式  用redis自带方法
	str, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Printf("conn.Do Set err : %v\n", err)
		return
	}
	fmt.Println("Get value = ", str)
}
