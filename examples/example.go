package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gophers-latam/GoKey/gokey"
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

	time.Sleep(10 * time.Second)
	res2, err := someGetOperation(client, "key")

	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(res2)
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
