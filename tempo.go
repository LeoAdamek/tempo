package tempo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

// APIHost is the host URL of the API
const APIHost = "https://api.tempo.io"

// Client represents a tempo Client
type Client struct {
	userAgent      string
	authToken      string
	tokenExpiresAt time.Time
	log            *log.Logger
	http           *http.Client

	// Enable this to log debugging information
	Debug bool
}

// New creates a new API client instance.
func New(token string) (*Client, error) {

	userAgent := fmt.Sprintf("tempo.go (%s %s %s)", runtime.GOOS, runtime.GOARCH, runtime.Version())

	return &Client{
		authToken:      token,
		tokenExpiresAt: time.Now().AddDate(1000, 0, 0),
		userAgent:      userAgent,
		log:            log.New(os.Stdout, "tempo", log.LstdFlags),
		http:           http.DefaultClient,
	}, nil
}

// Do will perform the given request.
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {

	token, err := c.RequestToken(ctx)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Authorization", "Bearer "+token)

	if c.Debug {
		req.Write(c.log.Writer())
		return nil, nil
	}

	res, err := ctxhttp.Do(ctx, c.http, req)

	if err != nil {
		c.log.Println("ERROR:", req.Method, req.URL.Path, err)
		return res, err
	}

	c.log.Println(req.Method, req.URL.Path, res.StatusCode, res.ContentLength)

	if res.StatusCode >= 400 {
		return res, getErrorFromResponse(res)
	}

	return res, nil
}

// JSON performs the given a request of the given method at the given path, encoding the given body as JSON
// and deserializing the JSON response into the given result interface
func (c *Client) JSON(ctx context.Context, method, path string, body, into interface{}) error {

	var b []byte
	buffer := bytes.NewBuffer(b)

	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return err
	}

	bodyReader := bytes.NewReader(buffer.Bytes())

	req, err := http.NewRequest(method, APIHost+"/core/3"+path, bodyReader)

	if err != nil {
		return err
	}

	res, err := c.Do(ctx, req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(into)
}
