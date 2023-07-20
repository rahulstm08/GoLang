package main
import "fmt"

type Employee struct{
	EmpId  int
	EmpName   string
	EmpMobile   []string
}

func main(){

	emp := Employee{
		EmpId: 101,
		EmpName : "Rahul",
		EmpMobile : []string{"62034950596"},
	}
	fmt.Println(emp)

		
}