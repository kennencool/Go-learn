package main

import "fmt"

func main(){
	var a = true
	var b = false
	if ( a && b ){
		fmt.Println("a, b 均为 true")
	}else if( a || b ){
		fmt.Println("a, b 有一个为 true")
	}else{
		fmt.Println("a, b 均为 false")
	}

	// 指针和地址操作符
	var c int = 4
	var ptr *int
	ptr = &c
	fmt.Println(c, ptr, *ptr)
}
