package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo // Create a new type called Todos, which is a slice of Todo structs.

// syntax for attaching a method on the slice
// func (receiverName ReceiverType) MethodName(parameters) returnType {
//     // method body
// }

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		fmt.Println(err)
		return err
	}

	return nil

}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	// Remove the todo at the specified index by slicing the slice.
	*todos = append(t[:index], t[index+1:]...) // Remove the todo at the specified index by combining the slices before and after the index.
	fmt.Printf("Todo at index %d deleted successfully.\n", index)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		complitionTime := time.Now()
		t[index].CompletedAt = &complitionTime

	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"

			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()

}
