package main

import (
	"fmt"
	"unsafe"
)
var a = "hello,world a"
var b string = "hello,world b"  // 这个string可以省略
var c bool  // 默认为false

var(
	x int
	y bool
)

var p, q = 1, 2
var m, n = "m", 1

// 常量
const c1, c2 = 123, "hahaha"
const(
	c3 = "123"
	c4 = len(c3)
	c5 = unsafe.Sizeof(c3)
)

// 常量枚举
const(
	UNKNOWN = 0
	FEMALE = 1
	MALE = 2
)

func main(){
	fmt.Println(a, b, c)
	fmt.Println(x, y)
	fmt.Println(p, q)
	fmt.Println(m, n)
	fmt.Println(c1, c2)
	fmt.Println(c3, c4, c5)
	fmt.Println(UNKNOWN, FEMALE, MALE)
}