package consts

import (
	"fmt"
	"strings"
)

const (
	// StatusTodo represents the "Todo" status
	StatusTodo = "Todo"
	// StatusInProgress represents the "In Progress" status
	StatusInProgress = "In Progress"

	// StatusDone represents the "Done" status
	StatusDone = "Done"

	// StatusNotStarted represents the "Not Started" status
	StatusNotStarted = "Not Started"

	// StatusOnHold represents the "On Hold" status
	StatusOnHold = "On Hold"

	// StatusCancelled represents the "Cancelled" status
	StatusCancelled = "Cancelled"

	// StatusBlocked represents the "Blocked" status
	StatusBlocked = "Blocked"
)

func GetAllStatuses() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
		StatusTodo,
		StatusInProgress,
		StatusDone,
		StatusNotStarted,
		StatusOnHold,
		StatusCancelled,
		StatusBlocked,
	)
}

func IsValidStatus(status string) bool {
	validStatuses := []string{
		StatusTodo,
		StatusInProgress,
		StatusDone,
		StatusNotStarted,
		StatusOnHold,
		StatusCancelled,
		StatusBlocked,
	}

	for _, validStatus := range validStatuses {
		if status == "" {
			return false
		}
		if strings.EqualFold(status, validStatus) {
			return true
		}
	}
	return false
}
