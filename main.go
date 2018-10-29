package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
)

var token = os.Getenv("TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

func main() {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	secret, err := c.Read("secret/data/foo")
	if err != nil {
		fmt.Println(err)
		return
	}
	m := secret.Data["data"].(map[string]interface{})
	fmt.Println(m["hello"])
}
