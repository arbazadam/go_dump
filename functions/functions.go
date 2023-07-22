package main

import "fmt"


type Result struct{
	 operation string
result int
}

//Your function can accept multiple arguments with the syntax below
func newFunc(operation string,values ...int) Result{
	result:=0
	switch (operation){
	case "addition":
		for _,val:=range values{
result=result+val
		}
	case "subtraction":
		for i,_:=range values{
			for j:=i+1;j<len(values);j++{
				if values[i]>values[j]{
if result>values[j]{
	result=result-values[j]
}else{
	result=values[i]-values[j]
}
				}
			}
					}
	default:
		
	}
finalResult:=Result{operation, result}
	return finalResult
}

func main(){
	// fmt.Println(newFunc("subtraction",4,3,4))
	a,b:=returnMultiple(10,9)
	fmt.Println(a,b)
	i:=Insaan{"Arbaz",20,true}
	fmt.Printf("The details of this insaan is %+v \n",i)
	var x *Insaan
	var y=func(){}
	fmt.Println(x)
	fmt.Println(y)
	//p := Person{Name: "John", Age: 30}

	// Additional properties stored in a map
	//Here when you make use of the the interface to define type of the values of a map, you basically mean that the values can be of any type
	additionalProperties := map[string]interface{}{
			"Address": "123 Main St",
			"test":10,
	}

	// Accessing the additional property
	//the ok variable is a type assertion which will return true if the Address key exists in the map.
	address, ok := additionalProperties["Address"]
	if ok {
			fmt.Println(address) // Output: 123 Main St
	}

	//fmt.Println(FindShortest([]int{10,9,232,12}))
}

type Person struct{
	Name string
	Age int
}

//if you have common type for two variables in your function definition, you can define them like this and also return them like the following.
func returnMultiple(a,b int) (x,y int){
	x=b
	y=a
	return
}

//There is no explicit inheritance with structs in golang, No super keyword as well
type Insaan struct{
	Name string
	age int
	status bool
}

func FindShortest(inputSlice []int) int{
	var shortest int
	for i:=0;i<len(inputSlice);i++{
		if i<len(inputSlice){
			if inputSlice[i]>inputSlice[i+1]{
				shortest=inputSlice[i+1]
			}
		}
	}
	return shortest
}