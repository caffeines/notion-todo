package notion

import (
	"errors"
	"fmt"
	"github.com/caffeines/notion-todo/service"
	"net/http"
)

// notionImpl is the implementation of Notion interface
type notionImpl struct {
	todoService       service.Todo
	credentialService service.Credential
}

var notion Notion

const URL = "https://api.notion.com/v1"

// NewNotionImpl returns a new instance of NotionImpl
func NewNotionImpl(todoSvc service.Todo, credService service.Credential) Notion {
	if notion == nil {
		notion = &notionImpl{
			todoService:       todoSvc,
			credentialService: credService,
		}
	}
	return notion
}

func (n *notionImpl) doRequest(req *http.Request) error {
	config, err := n.credentialService.GetConfig()
	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Notion-Version", "2022-06-28")

	// Send the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp)
		return errors.New("request failed")
	}

	fmt.Println("✅  Todo added successfully!")
	return nil
}

// AddPage adds a new page to the database
func (n *notionImpl) AddPage(title string) error {
	if n.todoService == nil {
		return errors.New("todo service is not initialized")
	}

	config, err := n.credentialService.GetConfig()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/pages", URL)
	body, err := n.todoService.GetCreateTodoBody(title, config.DatabaseID)
	if err != nil {
		return err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	err = n.doRequest(req)
	return err
}

// GetPages returns all the pages from the database
func (n *notionImpl) GetPages() error {
	url := fmt.Sprintf("%s/databases/%s/query", URL, "databaseID")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	err = n.doRequest(req)
	return err
}
