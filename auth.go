package tempo

import (
	"context"
	"time"
)

// RequestToken gets the request authorization token. Refreshing it if necessary
func (c *Client) RequestToken(ctx context.Context) (string, error) {
	if c.authToken == "" || c.tokenExpiresAt.Before(time.Now()) {
		if err := c.refreshToken(ctx); err != nil {
			return "", err
		}
	}

	return c.authToken, nil
}

func (c *Client) refreshToken(ctx context.Context) error {

	return nil
}
