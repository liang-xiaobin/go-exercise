package main

import "fmt"

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() *User {
	return &User{ID: 1888, Name: "liangxiaobin", Avatar: "www.liangxiaobin.com"}
}

func GetUserInfo1(u *User) *User {
	u.Avatar = "sdfsadf"
	return u
}

func GetUserInfo2(u User) *User {
	return &u
}

// main 测试变量内存是在栈或堆上分配
func main() {
	GetUserInfo()
	str := new(string)
	*str = "hello world"
	// 未确定类型 reflect.TypeOf(arg).Kind()
	fmt.Println(str)
	// 泄露参数
	GetUserInfo1(&User{ID: 12313})
	GetUserInfo2(User{ID: 234234})
}
