package main

import "fmt"

func main()  {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}

	for a < b {
		fmt.Println(a)
		a++
	}

	for i,item:= range numbers {
		fmt.Print(i)
		fmt.Print(":")
		fmt.Println(item)
	}
}
