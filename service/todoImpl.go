package service

import (
	"bytes"
	"encoding/json"
	"github.com/caffeines/notion-todo/models"
	"sync"
)

type TodoImpl struct{}

var (
	todo Todo
	once sync.Once
)

func NewTodoImpl() Todo {
	once.Do(func() {
		todo = &TodoImpl{}
	})
	return todo
}

// NewTodoProperties returns a new Properties
func (t *TodoImpl) newTodoProperties(title string) models.ItemData {
	return models.ItemData{
		ItemName: models.Title{
			Titles: []models.TextTitle{
				{
					Text: models.Text{
						Content: title,
					},
				},
			},
		},
		ItemStatus: models.Status{
			Select: models.Select{
				Name: "Todo",
			},
		},
	}
}

// GetCreateTodoBody returns a new Todo body
func (t *TodoImpl) GetCreateTodoBody(title string, databaseId string) (*bytes.Buffer, error) {
	itemData := t.newTodoProperties(title)
	payload := models.CreateTodoPayload{
		Parent: models.Parent{
			DatabaseID: databaseId,
		},
		Properties: itemData,
	}
	jsonPayload, err := json.Marshal(payload)
	return bytes.NewBuffer(jsonPayload), err
}
