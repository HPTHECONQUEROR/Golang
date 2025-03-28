package main
import "fmt"

type Employee struct{
	Name string
	Age int 
	Salary int
	Department string
}

func (e *Employee) DisplayDetails(){
	fmt.Println("Details of Employee ")
	fmt.Printf("Name: %s\nAge: %d\n,Salary: %d\nDepartment: %s",e.Name,e.Age,e.Salary,e.Department)
}

func (e *Employee) UpdateSalary(NewSalary int){
	e.Salary = NewSalary	
}

func main() {
	var employees []Employee
	var n int

	fmt.Print(employees)
	fmt.Print("Enter the nimber of employees: ")
	fmt.Scan(&n)

	for i:=1; i<n+1;i++{
		var temp Employee
		fmt.Println("Enter the details of employee (name,age,salary,dept): ")
		fmt.Scanln(&temp.Name,&temp.Age,&temp.Salary,&temp.Department)
		
		employees = append(employees, temp)
	}

	for _,val := range employees{
		val.DisplayDetails()	
	}
	
	employees[1].UpdateSalary(12323)
	employees[1].DisplayDetails()

	freq := make(map[string]int)

	for _,val := range employees{
		freq[val.Department] += 1
	}

	fmt.Println("freq: ",freq)

}
