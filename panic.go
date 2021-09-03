// package main

// import "fmt"

// func endApp() {
// 	// menangkap pesan error dari recover
// 	message:=recover()
// 	if message!=nil{

// 		fmt.Println("terjadi error dengan message:",message)
	

// 	}
// 	fmt.Println("Aplikasi selesai")
	
	

// }

// func runApplication(err bool)  {

// 	defer endApp()

// 	if err {
// 		// jika error dan berhenti ekskusi
// 		panic("Aplikasi error")

// 	}
// 	fmt.Println("aplikasi berjalan")


	
// }

// func main() {

// 	runApplication(false)


// }