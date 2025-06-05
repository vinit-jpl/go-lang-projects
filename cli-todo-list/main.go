package main

func main() {

	todos := Todos{} // initializing the Todos slice
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	// todos.add("Learn Go")
	// todos.add("Build a CLI Todo List")
	// todos.add("Test the application")

	// todos.toggle(0)
	// todos.print()
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)

}
