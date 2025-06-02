package main

import "fmt"

func main() {

	todos := Todos{}

	todos.add("Learn Go")
	todos.add("Build a CLI Todo List")
	todos.add("Test the application")

	fmt.Println("Current Todos:", todos)

	todos.delete(0)
	fmt.Println(todos)
}
