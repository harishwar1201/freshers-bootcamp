package main

import "fmt"
//hey
type Employee interface{
	GetSalary() int
}
type FulltimeEmployee struct {
	name string
	Duration int
}
type Contractor struct {
	name string
	Duration int
}
type Freelancer struct {
	name string
	Duration int
}
func (fulltimeEmp FulltimeEmployee) GetSalary() int {
	return fulltimeEmp.Duration*500
}
func (contr Contractor) GetSalary()int {
	return contr.Duration*100
}
func (flancer Freelancer) GetSalary() int {
	return flancer.Duration*10
}

func main(){
	employee1:=FulltimeEmployee{"Ramesh",28}
	employee2:=Contractor{"Rajesh",28}
	employee3:=Freelancer{"Suresh",20}
	fmt.Println("Total Salary of employee1 is ", employee1.GetSalary())
	fmt.Println("Total Salary of employee2 is ", employee2.GetSalary())
	fmt.Println("Total Salary of employee3 is ", employee3.GetSalary())
}


