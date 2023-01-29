package main

import "fmt"

/*
	type Employee struct {
	    Name         string
	    ID           string
	}

	func (e Employee) Description() string {
	    return fmt.Sprintf("%s (%s)", e.Name, e.ID)
	}

	type Manager struct {
	    Employee
	    Reports []Employee
	}

	func (m Manager) FindNewEmployees() []Employee {
	    // do business logic
	}
*/

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return m.Reports
}
func main() {
	/*

	   	m := Manager{
	          Employee: Employee{
	              Name:         "Bob Bobson",
	              ID:             "12345",
	          },
	          Reports: []Employee{},
	      }
	      fmt.Println(m.ID)            // prints 12345
	      fmt.Println(m.Description()) // prints Bob Bobson (12345)
	*/
	m := Manager{
		Employee: Employee{
			Name: "Hello",
			ID:   "123",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.Name, m.ID)

	/*
		type Inner struct {
		    X int
		}

		type Outer struct {
		    Inner
		    X int
		}
	*/
	type Inner struct {
		X int
	}
	type Outer struct {
		Inner
		X int
	}
	/*
		o := Outer{
		    Inner: Inner{
		        X: 10,
		    },
		    X: 20,
		}
		fmt.Println(o.X)       // prints 20
		fmt.Println(o.Inner.X) // prints 10
	*/
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X, o.Inner.X, o.Inner)

	//var eFail Employee = m        // compilation error!
	//var eOK Employee = m.Employee // ok!
	//var eFail Employee = m
	var eOk = m.Employee
	fmt.Println(eOk)

	/*
	 o := Outer{
	        Inner: Inner{
	            A: 10,
	        },
	        S: "Hello",
	    }
	    fmt.Println(o.Double())
	*/

	outer2 := Outer2{
		Inner2: Inner2{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(outer2.Double())
}

/*
type Inner struct {
    A int
}

func (i Inner) IntPrinter(val int) string {
    return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
    return i.IntPrinter(i.A * 2)
}

type Outer struct {
    Inner
    S string
}

func (o Outer) IntPrinter(val int) string {
    return fmt.Sprintf("Outer: %d", val)
}
*/

type Inner2 struct {
	A int
}

func (i Inner2) IntPrinter(val int) string {
	return fmt.Sprintf("Inner %d", val)
}
func (i Inner2) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer2 struct {
	Inner2
	S string
}

func (o Outer2) IntPrinter(val int) string {
	return fmt.Sprintf("Outer %d ", val)
}
