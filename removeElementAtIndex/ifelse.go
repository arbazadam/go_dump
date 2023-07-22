package main

import "fmt"

func main(){
	y:=removeFromSlice([]int{12,343,9,392},3)
	fmt.Println(y)
	if n:=10;n<15{
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Billa short")
	}
}

func removeFromSlice(slice []int,i int)[]int{
if slice!=nil && len(slice)>i && i>0{
newSlice:=append(slice[:i],slice[i+1:]...)
return newSlice
}
return slice
}