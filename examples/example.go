package main

import (
	"fmt"
	"github.com/gophers-latam/GoKey/gokey"
	"time"
)

func main() {
	client := gokey.NewClient()

	err := someSaveOperation(client, "key", "1")

	if err != nil {
		panic(err.Error())
	}

	res, err := someGetOperation(client, "key")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res)
}

func someSaveOperation(c *gokey.Client, key, value string) error {
	_, err := c.Save(key, []byte(value), time.Second*10)

	if err != nil {
		return err
	}
	return nil
}

func someGetOperation(client *gokey.Client, key string) (res string, err error) {
	b, err := client.Get(key)
	res = string(b)
	return res, err
}
