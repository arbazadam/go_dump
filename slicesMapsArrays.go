package main

import "fmt"

func createSliceArrayAndMaps(){
	fmt.Println("\nusing package")
	m:=[]int{1,23,4}
	fmt.Print(m)
}

func getLastThreeElements(n []int) {
	lastElementSlice := n[len(n)-1:]
	lastElement:=n[len(n)-1:]
	fmt.Println(lastElementSlice)
	fmt.Printf("the type of last element %T",lastElement) 
}