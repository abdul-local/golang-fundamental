// package main

// import (
// 	"container/ring"
// 	"fmt"
// 	"strconv"
// )


// func main() {

// 	data:=ring.New(5)

// 	for i:=0; i<data.Len(); i++{

// 		data.Value="Value ke-"+strconv.FormatInt(int64(i),10)
// 		data.Next()
// 		fmt.Println(data.Value)
// 	}

// }