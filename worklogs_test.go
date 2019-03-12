package tempo

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func TestClient_CreateWorklog(t *testing.T) {
	c, err := New("DEMO_TOKEN")

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	c.Debug = true

	w := Worklog{
		Start:        time.Date(2001, 11, 9, 8, 43, 10, 0, time.UTC),
		TimeSpent:    30 * time.Minute,
		BillableTime: 314159265 * time.Microsecond,
		UserID:       "some_random_id",
		Description:  "Time Description",
		IssueKey:     "TEST-1234",
	}

	if err := c.CreateWorklog(context.TODO(), &w); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestWorklog_MarshalJSON(t *testing.T) {

	expected := []byte(`{"issueKey":"TEST-1234","timeSpentSeconds":1800,"billableSeconds":315,"startDate":"2001-09-11","startTime":"08:43:08","authorAccountId":"some_random_id","description":"Time Description"}`)
	w := Worklog{
		Start:        time.Date(2001, 11, 9, 8, 43, 10, 0, time.UTC),
		TimeSpent:    30 * time.Minute,
		BillableTime: 314159265 * time.Microsecond,
		UserID:       "some_random_id",
		Description:  "Time Description",
		IssueKey:     "TEST-1234",
	}

	v, err := json.Marshal(w)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	for i, ex := range expected {
		if v[i] != ex {
			t.Errorf("Invalid character at position %d, expected %c but got %c", i, rune(ex), rune(v[i]))
			t.Fail()
		}
	}

	t.Logf("JSON Value: %s", v)
}
