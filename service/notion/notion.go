package notion

import "github.com/caffeines/notion-todo/models"

type Notion interface {
	AddPage(title, date string) error
	QueryPages(status, title string) ([]models.TodoItem, error)
	UpdatePageStatus(pageID, status string) error
	DeletePage(pageID string) error
}
