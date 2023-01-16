package test_redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

func TestHashSet(c redis.Conn) {
	for i := 0; i < 100; i++ {
		ss := "t_"
		field := ss + strconv.Itoa(i)
		data, err := redis.Int(c.Do("hset", "test_hash", field, 1))
		if err != nil {
			fmt.Println("redis err")
			return
		}
		fmt.Println(data)
	}
}
func TestGet(c redis.Conn) {
	data, err := redis.String(c.Do("Get", "test_str"))

	if err != nil {
		if err == redis.ErrNil {
			fmt.Println("data empty")
		} else {
			fmt.Println("redis错误", err)
		}
	}
	fmt.Println(data)
}
func TestSet(c redis.Conn) {
	for i := 0; i < 10000; i++ {
		// ss := "t_"
		// field := ss + strconv.Itoa(i)
		data, err := redis.Int(c.Do("sadd", "test_set", i))
		if err != nil {
			fmt.Println("redis err")
			return
		}
		fmt.Println(data)
	}

}
func TestHget(c redis.Conn) {
	data, err := redis.String(c.Do("hget", "test_hash", "ttt"))
	if err != nil {
		if err == redis.ErrNil {
			fmt.Println("not found this hash")
		} else {
			fmt.Printf("redis error", err.Error())
		}
	} else {
		println(data)
	}

}
func TestHscan(c redis.Conn) {

	var custor int
	custor = 0
	times := 0
	for {

		data, err := redis.Values(c.Do("hscan", "test_hash", custor))
		if err != nil {
			println("返回值错误", err)
		}
		data_info, _ := redis.Strings(data[1], err)
		for i := 0; i < len(data_info); i += 2 {
			println(data_info[i], data_info[i+1])
		}

		println(data_info, custor)
		//	for key, value := range date {
		//		println(key, value)
		//	}
		times += 1
		if custor == 0 {
			break
		}
	}
	println(times)

}
func TestRedisMo(command string) {

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis 错误")
		return
	}
	if c == nil {
		fmt.Println("connext is nil")
		return
	}
	defer c.Close()
	if command == "set" {
		test_redis.TestSet(c)
	}
	if command == "hscan" {
		TestHscan(c)
	}
	if command == "HashSet" {
		TestHashSet(c)
	}
	if command == "TestSet" {
		TestSet(c)
	}
	if command == "TestGet" {
		TestGet(c)
	}
	if command == "TestHget" {
		TestHget(c)
	}
}
