package main

import (
	"fmt"
	"strings"
)

func newFunc(){
  mySlice:=[]int{1,2,3,4}
  newRange:=mySlice[1:3]
  newRange=mySlice[1:] //from position one until the end
  newRange=mySlice[:3] //from start until 3 excluding 3
  mySlice=append(mySlice,10)//appends a value at the end of a slice
  fmt.Print(mySlice,newRange)
  fmt.Printf("the type is %T",newRange)
  fmt.Printf("%T",mySlice)

  myString:="arbaz"
  fmt.Println(strings.Contains(myString,"az"))
  fmt.Println(strings.ReplaceAll(myString,"az","dz"))

  //we can append to slices and ranges, but cannot append to an array
}