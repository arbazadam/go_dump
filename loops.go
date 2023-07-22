package main

import "fmt"

func checkLoops(){
for i:=10;i<20;i++{
	fmt.Println("This is my loop",i)
}
i:=5
for i<10{
	fmt.Println(i)
	i++
}
}