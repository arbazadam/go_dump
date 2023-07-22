package main

import (
	"fmt"
	"sort"
	"strings"
)

func mainReplit(){
  //shorthands cannot be used outside of the function
  //If we declare a variable and don't use it then its an error in goLang
  var name="arbaz"
  var newName string="roth"
  var testVar int
  shorthand:=10
  var decimal float64=12
  fmt.Println("New program",name,newName,testVar,shorthand,decimal)

  //Printf doesn't add new line
  fmt.Printf("printing something else %v",newName)

  //we also have Sprintf it basically saves the formatted string in a variable.
  //Sprintf doesnt print the value to the console it returns us the value.

  var abc=fmt.Sprintf("this is a formatted string %T \n",testVar)
  fmt.Print(abc)

  test:="hi this is from golang"
  fmt.Println(strings.ReplaceAll(test,"is","are")) //returns new string
  fmt.Println(test)

  //creating arrays in go
  var myArray [2]string=[2]string{"plus","minus"}
  //array shorthand
  var mySecondArray=[3]string{"make","model"}
  var mySlice=[]int{600,40,50}

  xyz:=sort.SearchInts(mySlice,40)
  fmt.Println(xyz)

  //sort.Ints(mySlice)
  // fmt.Println(mySlice)
  //in golang the for keyword is used for all types of loops, for,while or for in
   fmt.Println(myArray,mySecondArray)

  myVar:=0

  for myVar<4{ //this is a while loop
    fmt.Println("the val is",myVar)
    myVar++
  }

  for i:=1;i<10;i++{
    fmt.Println("the value inside loop is",i) //traditional for loop
  }
  rangeSlice(mySlice)
  printGreeting([]string{"arbaz","rahul","ramesh","adil"},greetUser)
}

func greetUser(name string){
  fmt.Print("Hello ",name,"\n")
}

func myLoopSlice(myTestSlice []int){
  for i:=0;i<len(myTestSlice);i++{
    fmt.Println("the slice values are: ",myTestSlice[i])
  }
}

func rangeSlice(mySlice []int){
  for index,value:=range mySlice{
    fmt.Printf("the position at index %v %v \n",index,value)
  }
}

func printGreeting(nameSlice[] string,f func(string)){
  for _,val:=range nameSlice{
    f(val)
  }
}