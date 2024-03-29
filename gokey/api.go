package gokey

import (
	"time"
)

type Client struct {
	cache Operations
}

func NewClient() *Client {
	return newClient()
}

func newClient() *Client {
	return &Client{
		cache: newCache(),
	}
}

func (c *Client) Save(k string, v []byte, ttl time.Duration) (bool, error) {
	return c.cache.Upsert(k, v, ttl)
}

func (c *Client) Get(key string) ([]byte, error) {
	return c.cache.Get(key)
}

func (c *Client) Delete(key string) (bool, error) {
	return c.cache.Delete(key)
}
