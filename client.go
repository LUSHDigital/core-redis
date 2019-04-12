package coreredis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Client represents a wrapped redis client.
type Client struct {
	*redis.Client
}

// Check returns information about the connection to the redis server.
func (c *Client) Check() ([]string, bool) {
	status := c.Ping()
	res, err := status.Result()
	if err != nil {
		msgs := []string{"redis did not respond to PING"}
		if res != "" {
			msgs = append(msgs, res)
		}
		return msgs, false
	}
	return []string{fmt.Sprintf("redis responded to PING with: %s", res)}, true
}
