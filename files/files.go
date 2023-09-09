package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
scanner:=bufio.NewScanner(os.Stdin)
fmt.Println("Please enter content to print")
scanner.Scan()
scanner.Text()
fmt.Println(convertByteToString([]byte{67,70,20,13}))
//this method creates a new file and writes to it if the file doesnt exist. If the file exists and contains data, the data will be truncated
//os.WriteFile("output.txt", []byte(x), 0644)
}

func convertByteToString(slice []byte) string{
	if len(slice)>0{
 x:=string(slice)
 return x;
	}
	return "";
}