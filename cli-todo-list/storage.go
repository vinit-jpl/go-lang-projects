package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct { // T can be of any type, in our case T is of type Todos
	Filename string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{Filename: filename}
}

/*
// below is a method called save that belongs to a pointer to a generic storage struct

	func(receiverName ReceiverType) methodName(params) returnType {
		 method body
	}
*/
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
