package tempo

import (
	"fmt"
	"net/http"
)

// Error represents an error returned by the Tempo API
type Error struct {
	responseBody []byte
	Code         string
}

func (e Error) Error() string {
	return fmt.Sprintf("Tempo Error(%s): %s", e.Code, "")
}

func getErrorFromResponse(res *http.Response) Error {
	e := Error{
		Code: res.Status,
	}

	return e
}
