package main

import "fmt"

func main(){

	// empSal := make(map[string]int)  //declaration

	// empSal = map[string]int{       //intialsation
	// 	"rahul" : 1000,
	// 	"atul" : 500,
	// 	"ankit" : 2000,
	// }


	// empSal := map[string]int{       //declaration with intialsation
	// 	"rahul" : 1000,
	// 	"atul" : 500,
	// 	"ankit" : 2000,
	// }
	// empSal["manish"] = 1500

	// fmt.Println(empSal)


	//delete method

	// empSal := map[string]int{       
	// 	"rahul" : 1000,
	// 	"atul" : 500,
	// 	"ankit" : 2000,
	// }
	
	// delete(empSal,"atul")
	// fmt.Println(empSal)

	//avalibility check
	empSal := map[string]int{       //declaration with intialsation
		"rahul" : 1000,
		"atul" : 500,
		"ankit" : 2000,
	}
	
	// ankit,ok := empSal["ankit"]
	// fmt.Println(ankit,ok)

	neha,ok := empSal["neha"]
	fmt.Println(neha,ok)

}