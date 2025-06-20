package notion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/caffeines/notion-todo/service/config"
	"github.com/caffeines/notion-todo/service/utility"

	"github.com/caffeines/notion-todo/consts"
	"github.com/caffeines/notion-todo/models"
)

// notionImpl is the implementation of Notion interface
type notionImpl struct {
	credentialService config.Credential
}

var notion Notion

// NewNotionImpl returns a new instance of NotionImpl
func NewNotionImpl(credService config.Credential) Notion {
	if notion == nil {
		notion = &notionImpl{
			credentialService: credService,
		}
	}
	return notion
}

func (n *notionImpl) doRequest(req *http.Request) error {
	config, err := n.credentialService.GetConfig()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return err
	}
	// Set the necessary headers
	req.Header.Set("Content-Type", consts.CONTENT_TYPE)
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Notion-Version", consts.NOTION_VERSION)

	// Send the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read response body for better error reporting
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return fmt.Errorf("failed to read response body: %v", readErr)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notion API error (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}

// AddPage adds a new page to the database
func (n *notionImpl) AddPage(title, date string) error {
	config, err := n.credentialService.GetConfig()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/pages", consts.API_URL)
	body, err := utility.GetCreateTodoBody(title, date, config.DatabaseID)
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
	url := fmt.Sprintf("%s/databases/%s/query", consts.API_URL, "databaseID")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	err = n.doRequest(req)
	return err
}

// QueryPages queries pages from the Notion database with optional filters
func (n *notionImpl) QueryPages(status, title string) ([]models.TodoItem, error) {
	config, err := n.credentialService.GetConfig()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/databases/%s/query", consts.API_URL, config.DatabaseID)

	// Build query filter
	var queryReq models.QueryRequest
	var filters []map[string]interface{}

	// Add status filter if provided
	if status != "" {
		statusFilter := map[string]interface{}{
			"property": "Status",
			"select": map[string]interface{}{
				"equals": status,
			},
		}
		filters = append(filters, statusFilter)
	}

	// Add title filter if provided
	if title != "" {
		titleFilter := map[string]interface{}{
			"property": "Title",
			"title": map[string]interface{}{
				"contains": title,
			},
		}
		filters = append(filters, titleFilter)
	}

	// Only add filter if we have conditions
	if len(filters) > 0 {
		queryReq.Filter = &models.QueryFilter{
			And: filters,
		}
	}

	// Convert to JSON
	jsonData, err := json.Marshal(queryReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	config, err = n.credentialService.GetConfig()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", consts.CONTENT_TYPE)
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Notion-Version", consts.NOTION_VERSION)

	// Send request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("notion API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var queryResp models.NotionQueryResponse
	if err := json.Unmarshal(body, &queryResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Convert to TodoItems
	var todos []models.TodoItem
	for _, page := range queryResp.Results {
		todos = append(todos, page.ToTodoItem())
	}

	return todos, nil
}

// UpdatePageStatus updates the status of a specific page in Notion
func (n *notionImpl) UpdatePageStatus(pageID, status string) error {
	if n.credentialService == nil {
		return errors.New("credential service is not initialized")
	}

	config, err := n.credentialService.GetConfig()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/pages/%s", consts.API_URL, pageID)

	// Create request body for status update
	updateReq := map[string]interface{}{
		"properties": map[string]interface{}{
			"Status": map[string]interface{}{
				"select": map[string]interface{}{
					"name": status,
				},
			},
		},
	}

	// Convert to JSON
	jsonData, err := json.Marshal(updateReq)
	if err != nil {
		return fmt.Errorf("failed to marshal update request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", consts.CONTENT_TYPE)
	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Notion-Version", consts.NOTION_VERSION)

	// Send request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body for error reporting
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("notion API error (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}
