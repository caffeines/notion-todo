package service

import "bytes"

type Todo interface {
	GetCreateTodoBody(title string, databaseId string) (*bytes.Buffer, error)
}
