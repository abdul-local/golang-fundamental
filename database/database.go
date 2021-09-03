package database

import "fmt"


var connection string



func init() {
	fmt.Println("blank idefier")

	connection="MYSQL"
	
}

func GetDatabase() string  {

	return connection
	
}