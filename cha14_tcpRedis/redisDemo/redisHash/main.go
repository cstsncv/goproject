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

	//2. 通过Go向redis写入数据 Hash
	_, err = conn.Do("HSet", "user1", "name", "Tom猫") //hset写一对儿
	if err != nil {
		fmt.Printf("conn.Do HSet err : %v\n", err)
		return
	}
	_, err = conn.Do("HMSet", "user1", "age", 11, "addr", "BeiJing", "score", 55.2) //hmset 写多对儿
	if err != nil {
		fmt.Printf("conn.Do HMSet err : %v\n", err)
		return
	}
	fmt.Println("HSet OK")
	//3. 通过Go向redis读取数据 string
	//返回的是interface{}空接口,需要转换成对应格式  用redis自带方法
	str, err := redis.String(conn.Do("HGet", "user1", "name")) //取一个
	if err != nil {
		fmt.Printf("conn.Do HSet err : %v\n", err)
		return
	}
	fmt.Println("user1.name = ", str)

	strs, err := redis.Strings(conn.Do("HMGet", "user1", "name", "age", "addr", "score")) //取一堆 取回为[]string
	if err != nil {
		fmt.Printf("conn.Do Het err : %v\n", err)
		return
	}
	fmt.Printf("HGet value = %v, strs type = %T\n", strs, strs)
	for i, v := range strs {
		fmt.Printf("user1.%v = %v\n", i, v)
	}
}
