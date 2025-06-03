package main

func main() {

	todos := Todos{}

	todos.add("Learn Go")
	todos.add("Build a CLI Todo List")
	todos.add("Test the application")

	todos.toggle(0)
	todos.print()

}
