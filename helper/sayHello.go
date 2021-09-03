package helper 

import "fmt"


// ini tidak bisa di akses dari luar package
var version="1.0.1"

// ini bisa di akses dari luar karena diawali dengan huruf besar
var Application="Belajar Golang"

func SayHello() {

	fmt.Println("Hay mas bro")
	
}