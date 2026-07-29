package main

import (
	//"bufio"
	"fmt"
	//"os"
	////"strconv"
	//"strings"
)

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	/* r := bufio.NewReader(os.Stdin)

	// ID
	fmt.Print("Enter ID: ")
	idStr, _ := r.ReadString('\n')

	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("The ID should be a number")
		return
	}
	e.ID = id

	// Name
	fmt.Print("Enter name: ")
	name, _ := r.ReadString('\n')
	e.Name = strings.TrimSpace(name)

	// Age
	fmt.Print("Enter age: ")
	ageStr, _ := r.ReadString('\n')

	age, err := strconv.Atoi(strings.TrimSpace(ageStr))
	if err != nil {
		fmt.Println("The age should be a number")
		return
	}
	e.Age = age

	// Salary
	fmt.Print("Enter salary: ")
	salaryStr, _ := r.ReadString('\n')

	salary, err := strconv.ParseFloat(strings.TrimSpace(salaryStr), 64)
	if err != nil {
		fmt.Println("The salary should be a number")
		return
	}
	e.Salary = salary

	// Add employee to the manager's slice
	m.Employees = append(m.Employees, e) */
	m.Employees = append(m.Employees, e)
	
	
	
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
	// TODO: Implement this method
	for i:=0; i<len(m.Employees);i++{
	    if m.Employees[i].ID == id{
	        m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
	    }
	}
	    
	}
	


// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	var avr float64 = 0
	if len(m.Employees) == 0 {
	    return 0
	}
	for i:=0; i<len(m.Employees); i++{
		avr +=m.Employees[i].Salary
	}
	return avr/float64(len(m.Employees))
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	// TODO: Implement this method
	for i:=0; i<len(m.Employees);i++{
	    if m.Employees[i].ID == id{
	        return &m.Employees[i]
	    }
	}
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
