package main

import "fmt"

// Encapsulation: Defining a struct with fields and methods
type Person struct {
	Name string
	Age  int
}

// Method for the Person struct (Behavior)
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Composition: Employee "has a" Person (embedding Person struct)
type Employee struct {
	Person
	Position string
	Salary   int
}

// Method for Employee (Behavior specific to Employee)
func (e Employee) Work() {
	fmt.Printf("%s is working (age %d) as a %s.\n", e.Name, e.Age, e.Position)
}

// Polymorphism: Interface for actions that can be performed
type Worker interface {
	Work()
}

func DoWork(w Worker) {
	w.Work() // Calls the Work method of the struct that implements the Worker interface
}

func main() {
	// Create a Person instance
	person := Person{Name: "Alice", Age: 30}
	person.Greet() // Encapsulation: Greet method is bound to Person

	// Create an Employee instance using Composition (Employee has a Person)
	employee := Employee{
		Person:   Person{Name: "Bob", Age: 40},
		Position: "Engineer",
		Salary:   50000,
	}

	// Access embedded fields and methods
	employee.Greet() // Employee can access Person's method
	employee.Work()  // Employee's specific method

	// Polymorphism: DoWork can accept any type that implements the Worker interface
	DoWork(employee) // Employee implements the Worker interface
}
