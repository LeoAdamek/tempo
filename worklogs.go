package tempo

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"time"
)

// Worklog represents a single worklog entry
type Worklog struct {
	IssueKey     string
	TimeSpent    time.Duration
	BillableTime time.Duration
	Start        time.Time
	UserID       string
	Description  string
}

// CreateWorklog will create the given worklog
func (c *Client) CreateWorklog(ctx context.Context, worklog *Worklog) error {
	return c.JSON(ctx, http.MethodPost, "/worklogs", worklog, &worklog)
}

// MarshalJSON marshals the Worklog into JSON
func (w Worklog) MarshalJSON() ([]byte, error) {
	timeSpentSeconds := int64(math.Ceil(w.TimeSpent.Seconds()))
	// billableTimeSeconds := int64(math.Ceil(w.BillableTime.Seconds()))
	startDate := w.Start.Format("2006-02-01")
	startTime := w.Start.Format("15:04:03")

	return json.Marshal(&struct {
		IssueKey         string `json:"issueKey"`
		TimeSpentSeconds int64  `json:"timeSpentSeconds"`
		// BillableSeconds  *int64 `json:"billableSeconds,omitempty"`
		StartDate       string `json:"startDate"`
		StartTime       string `json:"startTime"`
		AuthorAccountID string `json:"authorAccountId"`
		Description     string `json:"description"`
	}{
		IssueKey:         w.IssueKey,
		TimeSpentSeconds: timeSpentSeconds,
		// BillableSeconds:  billableTimeSeconds,
		StartDate:       startDate,
		StartTime:       startTime,
		AuthorAccountID: w.UserID,
		Description:     w.Description,
	})
}
