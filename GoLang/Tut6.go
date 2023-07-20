package main

import "fmt"

func main(){

	
    //array
	// var arr [4]int = [4]int{1,2,3,4}
	// var arr = [...]{1,2,3,4}

	//slicing
	// var arr []int = []int{1,2,3,4}
	// fmt.Println(len(arr))

	// add enew element
	var arr []int = []int{1,2,3,4}
	arr = append(arr, 100)

	fmt.Println(arr)

}