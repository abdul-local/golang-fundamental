// package main

// import "fmt"

// type Address struct {
// 	City, Province, Country string
// }

// func main() {

// 	var address1 Address = Address{
 
// 		City:     "Cirebon",
// 		Province: "Bandung",
// 		Country:  "Indonesia",
// 	}

// 	var address2 *Address= &address1
// 	var address3 *Address = &address1

	// pass by value
	// address2 := address1
	// address2.City = "Subang"

	

	// fmt.Println("address1",address1)
	// fmt.Println("address2",address2)


	// pass by refrence

	// address2 := &address1
	// address2.City = "Subang"



    // yang ini seolah-olah kita buat yang struct yang baru tapi tidak merubah yang lama
	// address2 = &Address{

	// 	City:     "pulau seribu",
	// 	Province: "DKI-Jakrata",
	// 	Country:  "Indonesia",
	// }

	// membuat yang baru dan merubah yang lama

// 	*address2 = Address{

// 		City:     "pulau seribu",
// 		Province: "DKI-Jakrata",
// 		Country:  "Indonesia",
// 	}


	
// 	fmt.Println("address1",address1)
// 	fmt.Println("address2",address2)
// 	fmt.Println("address3",address3)

// 	var address4 *Address=new(Address)
// 	address4.City="mataram"
// 	address4.Province="NTB"
// 	address4.Country="IDN"

// 	fmt.Println(address4)


// }