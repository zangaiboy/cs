package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

//自定义的redis实例的端口为32771
func main() {
	//使用redis封装的Dial进行tcp连接
	c, err := redis.Dial("tcp", "192.168.86.128:6379")
	errCheck(err)

	defer c.Close()

	//对本次连接进行set操作
	_, setErr := c.Do("set", "url", "xxbandy.github.io")
	errCheck(setErr)

	//使用redis的string类型获取set的k/v信息
	r, getErr := redis.String(c.Do("get", "url"))
	errCheck(getErr)
	fmt.Println(r)
}
func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
