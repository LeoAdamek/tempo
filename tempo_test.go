package tempo

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_JSON(t *testing.T) {

	body := &struct {
		String string `json:"string"`
	}{
		String: "some string",
	}

	c, err := New("DEMO_TOKEN")

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	c.Debug = true

	ctx := context.TODO()

	ma := make(map[string]interface{})

	if err := c.JSON(ctx, http.MethodPost, "/test", body, &ma); err != nil {
		// t.Error("JSON Error:", err)
		// t.Fail()
	}

}
