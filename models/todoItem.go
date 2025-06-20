package models

import (
	"github.com/caffeines/notion-todo/consts"
)

type Text struct {
	Content string `json:"content"`
}

type TextTitle struct {
	Text Text `json:"text"`
}

type Title struct {
	Titles []TextTitle `json:"title"`
}
type Select struct {
	Name string `json:"name"`
}
type Status struct {
	Select Select `json:"select"`
}

type DateValue struct {
	Start *string `json:"start,omitempty"`
}

type Date struct {
	Value DateValue `json:"date"`
}

type ItemData struct {
	ItemName   Title  `json:"Title"`
	ItemStatus Status `json:"Status"`
	ItemDate   *Date  `json:"Due Date,omitempty"`
}

type Properties struct {
	Item ItemData `json:"properties"`
}

// NewProperties returns a new Properties
func NewProperties(title, date string) ItemData {
	data := ItemData{
		ItemName: Title{
			Titles: []TextTitle{
				{
					Text: Text{
						Content: title,
					},
				},
			},
		},
		ItemStatus: Status{
			Select: Select{
				Name: consts.StatusTodo,
			},
		},
		ItemDate: nil,
	}

	if date != "" {
		data.ItemDate = &Date{
			Value: DateValue{
				Start: &date,
			},
		}
	}
	return data
}

// Response models for Notion query API
type NotionTextContent struct {
	Type string `json:"type"`
	Text struct {
		Content string  `json:"content"`
		Link    *string `json:"link"`
	} `json:"text"`
	PlainText string `json:"plain_text"`
}

type NotionTitleProperty struct {
	ID    string              `json:"id"`
	Type  string              `json:"type"`
	Title []NotionTextContent `json:"title"`
}

type NotionSelectOption struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type NotionSelectProperty struct {
	ID     string              `json:"id"`
	Type   string              `json:"type"`
	Select *NotionSelectOption `json:"select"`
}

type NotionDateValue struct {
	Start    *string `json:"start"`
	End      *string `json:"end"`
	TimeZone *string `json:"time_zone"`
}

type NotionDateProperty struct {
	ID   string           `json:"id"`
	Type string           `json:"type"`
	Date *NotionDateValue `json:"date"`
}

type NotionPageProperties struct {
	Title   NotionTitleProperty  `json:"Title"`
	Status  NotionSelectProperty `json:"Status"`
	DueDate NotionDateProperty   `json:"Due Date"`
}

type NotionPage struct {
	Object         string               `json:"object"`
	ID             string               `json:"id"`
	CreatedTime    string               `json:"created_time"`
	LastEditedTime string               `json:"last_edited_time"`
	Properties     NotionPageProperties `json:"properties"`
	URL            string               `json:"url"`
}

type NotionQueryResponse struct {
	Object     string       `json:"object"`
	Results    []NotionPage `json:"results"`
	NextCursor *string      `json:"next_cursor"`
	HasMore    bool         `json:"has_more"`
}

// TodoItem represents a simplified todo item from Notion
type TodoItem struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Status  string  `json:"status"`
	DueDate *string `json:"due_date"`
	URL     string  `json:"url"`
}

// Convert NotionPage to TodoItem
func (p *NotionPage) ToTodoItem() TodoItem {
	item := TodoItem{
		ID:  p.ID,
		URL: p.URL,
	}

	// Extract title
	if len(p.Properties.Title.Title) > 0 {
		item.Title = p.Properties.Title.Title[0].PlainText
	}

	// Extract status
	if p.Properties.Status.Select != nil {
		item.Status = p.Properties.Status.Select.Name
	}

	// Extract due date
	if p.Properties.DueDate.Date != nil && p.Properties.DueDate.Date.Start != nil {
		item.DueDate = p.Properties.DueDate.Date.Start
	}

	return item
}

// QueryFilter represents the filter for Notion query
type QueryFilter struct {
	And []map[string]interface{} `json:"and,omitempty"`
}

type QueryRequest struct {
	Filter *QueryFilter `json:"filter,omitempty"`
}
