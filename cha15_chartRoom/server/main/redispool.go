package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

//初始化redis连接池
var pool *redis.Pool

func initRedis(addressAndPort string, protocol string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   //和数据库最大连接数,0表示无限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接匿名函数,用于连接redis server,返回连接
			return redis.Dial(protocol, addressAndPort)
		},
	}
}
