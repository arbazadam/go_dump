package main

import "fmt"


type Result struct{
	 operation string
result int
}

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
	fmt.Print(a,b)
}

func returnMultiple(a,b int) (x,y int){
	x=b
	y=a
	return
}