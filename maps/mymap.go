package main

import "fmt"

func main(){
	//there are two ways to create a map, one using the map literal and the other way is using the make function.

	literalMap:=map[int]string{
		1:"arbaz",
		2:"adil",
	}

	otherMap:=make(map[string]string)

	otherMap["1"]="Java"
	otherMap["2"]="Python"
	otherMap["3"]="Javascript"

	delete(otherMap,"1")

	fmt.Println(literalMap)
	fmt.Println(otherMap)
}