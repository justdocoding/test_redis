package test_redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
)

func test_set(c net.Conn) {
	c.Do("Set", "test", 123)
	e, _ := redis.Int(c.Do("Get", "test"))
	fmt.Println(e)
}
