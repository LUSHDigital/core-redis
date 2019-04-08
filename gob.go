package redis

import (
	"bytes"
	"encoding/gob"
	"time"
)

// NewGobClient returns a redis client wrapper for Gob operations.
func NewGobClient(c *Client) *GobClient {
	return &GobClient{c}
}

// GobClient represents a wrapped redis client for Gob operations.
type GobClient struct {
	*Client
}

// GetGob will get try and get byte data from a key in redis and gob unmarshal to the provided container.
func (c *GobClient) GetGob(key string, v interface{}) error {
	var data []byte
	data, err := c.Get(key).Bytes()
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(data)
	return gob.NewDecoder(b).Decode(v)
}

// SetGob will take the value and gob marshal it to bytes and store in redis.
func (c *GobClient) SetGob(key string, v interface{}, expiration time.Duration) error {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(v); err != nil {
		return err
	}
	data := b.Bytes()
	return c.Set(key, data, expiration).Err()
}
