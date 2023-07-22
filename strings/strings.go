package main

import "fmt"

func main(){
	x:="my name is arbaz"
	for i,val:=range x{
		fmt.Printf("Character at index %v: %c\n",i,val)
	}
}