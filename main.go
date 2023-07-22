package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a number")
	scanner.Scan()
	x:=scanner.Text()
	p,_:=strconv.Atoi(x)
	y:=reverseTheNumber(p)
	fmt.Println(y)
	//checkLoops()
	j:=learnPointer()
	fmt.Println(*j)
	t:=*j
	t="rizwan"
	fmt.Println(&t)
	fmt.Println(t)
	fmt.Println(*j)
}

func reverseTheNumber(x int) int{
	return 10
}