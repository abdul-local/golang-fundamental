// package main

// import (
// 	"fmt"
	
// )

// type Address struct {
// 	City,Province, Country string
// }


// func main() {

	

// 	// address1:=Address{

// 	// 	City: "Jakarta",
// 	// 	Province: "DKI-Jakarta",
// 	// 	Country: "IDN",

// 	// }
// 	 var address1 Address=Address{

// 		City: "Jakarta",
// 		Province: "DKI-Jakarta",
// 		Country: "IDN",

// 	}
// 	var address2 *Address=&address1
// 	var address3 *Address=&address1

// 	var address4 *Address = new(Address)

	


// 	// pass by value (beneran di copy datanya)
// 	// address2 := address1
// 	// address2.City="Tanah Abang"
	

// 	// fmt.Println(address1)
// 	// fmt.Println(address2)

// 	// pass by refrence
//     // & merujuk ke reference (pointer)
// 	// address2:=&address1
// 	// ini di rubah ikut juga berubah di addrees1
// 	// address2.City="Tanah abang"

// 	*address2 =Address{

// 		City: "Lombok",
// 		Province: "NTB",
// 		Country: "IDN",

// 	}



// 	fmt.Println(address1)

// 	fmt.Println(address2)
// 	fmt.Println(address3)
// 	fmt.Println(address4)




	


	


// }