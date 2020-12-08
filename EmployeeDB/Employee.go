package main

import (
"fmt"
)

func main() {

	fmt.Println("Printing Employee(s) in DB...")
	fmt.Println()
	
	createEmp()
	
	fmt.Println()

	str1 := "Done "
	str2 := "Printing!"
	fmt.Println(str1 + str2)
}

type employee struct {
	firstName string
	lastName string
	jobCode int
	title string
	salary int
}

func (x employee) printEmp(){
	fmt.Println("Emp first name: ", x.firstName)
	fmt.Println("Emp last name: ", x.lastName)
	fmt.Println("Job Code: ", x.jobCode)
	fmt.Println("Title: ", x.title)
	fmt.Println("salary: ", x.salary)
}

func createEmp(){
	testEmp := employee{
		firstName: "John",
		lastName: "Doe",
		jobCode: 439024,
		title: "Software Engineer",
		salary: 75000,
	}
	testEmp.printEmp()
}



