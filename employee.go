package employee

// Employee represents an employee with name and age.
type Employee struct {
	name string
	age  int
}

// NewEmployee creates a new instance of Employee.
func NewEmployee(name string, age int) *Employee {
	return &Employee{
		name: name,
		age:  age,
	}
}

// GetName returns the name of the employee.
func (e *Employee) GetName() string {
	return e.name
}

// GetAge returns the age of the employee.
func (e *Employee) GetAge() int {
	return e.age
}
