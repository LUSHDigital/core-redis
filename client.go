package coreredis

import (
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
		return []string{res}, false
	}
	return []string{res}, true
}
