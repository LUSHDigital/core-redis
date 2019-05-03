package coreredis

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

const (
	// DefaultURL is the default redis host url.
	DefaultURL = "redis://127.0.0.1:6379/0"
)

// ConfigError happens when the redis connection cannot be configured.
type ConfigError struct {
	msg interface{}
}

func (e ConfigError) Error() string {
	return fmt.Sprintf("redis config error: %v", e.msg)
}

// URLFromEnv tries to retrieve the redis url from the environment.
func URLFromEnv() string {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		url = os.Getenv("REDIS_HOST")
	}
	if url == "" {
		url = DefaultURL
	}
	return url
}

// Parse attempts to parse a redis url and return options.
func Parse(s string) (*redis.Options, error) {
	rurl, err := url.Parse(s)
	if err != nil {
		return nil, ConfigError{err}
	}
	pass, ok := rurl.User.Password()
	if !ok {
		pass = rurl.User.Username()
	}

	var db int
	path := strings.Split(strings.TrimPrefix(rurl.Path, "/"), "/")
	if len(path) > 1 {
		if path[0] != "" {
			n, err := strconv.Atoi(path[0])
			if err != nil {
				return nil, ConfigError{err}
			}
			db = n
		}
	}

	opt := &redis.Options{
		Addr:     rurl.Host,
		Password: pass,
		DB:       db,
	}
	return opt, nil
}

// NewDefaultClient returns a wrapped redis client with default configuration.
func NewDefaultClient() *Client {
	opt, err := Parse(URLFromEnv())
	if err != nil {
		log.Fatalln(err)
	}
	return NewClient(opt)
}

// NewClient returns a wrapped redis client.
func NewClient(opt *redis.Options) *Client {
	return &Client{redis.NewClient(opt)}
}
