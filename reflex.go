// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Sample struct {
// 	Name string ` required:"true" max:"10"`
// }

// type ContohLagi struct {
// 	Name string ` required:"true" max:"10"`
// 	Description string
// }

// func IsValid(data interface{}) bool  {
// 	t:=reflect.TypeOf(data)
// 	for i:=0; i<t.NumField(); i++ {
// 		field:=t.Field(i)
// 		if field.Tag.Get("required") == "true" {

// 			value:=reflect.ValueOf(data).Field(i).Interface()

// 			if value ==""{
// 				return false
// 			}

// 		}

		
// 	}

// 	return true
	
// }


// func main() {

// 	sample:=Sample{"abdul"}
// 	sampleType:=reflect.TypeOf(sample)
// 	structurField:=sampleType.Field(0)
// 	required:=structurField.Tag.Get("required")
// 	max:=structurField.Tag.Get("max")
// 	fmt.Println(structurField)
// 	fmt.Println(structurField.Name)
// 	fmt.Println(required)
// 	fmt.Println(max)

// 	sample.Name=""
// 	contoh:=ContohLagi{"abdul","hamzan"}

// 	fmt.Println(IsValid(sample))
// 	fmt.Println(IsValid(contoh))
	

// }



