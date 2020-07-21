package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

const (
	VAULT_ADDR = ""
	TOKEN      = ""
)

func main() {

	client, err := api.NewClient(&api.Config{
		Address: VAULT_ADDR,
	})

	if err != nil {
		log.Fatal(err)
	}

	client.SetToken(TOKEN)

	data := map[string]interface{}{
		"data": map[string]string{
			"foo": "bar",
			"zip": "zap",
		},
	}

	client.Logical().Write("secret/data/my-secret", data)

	mydata, err := client.Logical().Read("secret/data/my-secret")
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(mydata.Data)
	fmt.Println(string(b))
	tk, err := mydata.TokenID()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(tk)

	s, err := client.Logical().List("auth/token/accessors")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", s)

	data = map[string]interface{}{
		"meta": map[string]string{
			"user": "armon",
		},
		"ttl":       "1h",
		"renewable": true,
	}
	mydata, err = client.Logical().Write("auth/token/create", data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", mydata)
	b, err = json.Marshal(mydata)
	if err != nil {
		panic(err)
	}
	fmt.Println("auth response", string(b))

	mydata, err = client.Logical().List("auth/token/roles")
	if err != nil {
		panic(err)
	}

	b, _ = json.Marshal(mydata)
	fmt.Println(string(b))
}
