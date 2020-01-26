package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-chef/chef"
)

func main() {
	// read a client key
	key, err := ioutil.ReadFile("george.pem")
	if err != nil {
		fmt.Println("Couldn't read key.pem:", err)
		os.Exit(1)
	}

	// build a client
	client, err := chef.NewClient(&chef.Config{
		Name: "george",
		Key:  string(key),
		// goiardi is on port 4545 by default. chef-zero is 8889
		BaseURL: "https://chef.senao.com.tw/organizations/senao/",
	})
	if err != nil {
		fmt.Println("Issue setting up client:", err)
	}

	// List Cookbooks
	cookList, err := client.Cookbooks.List()
	if err != nil {
		fmt.Println("Issue listing cookbooks:", err)
	}

	// Print out the list
	fmt.Println(cookList)


	nodelist, err := client.Nodes.List()
	if err != nil {
		fmt.Println("Issue listing Nodes:", err)
	}
	fmt.Println(nodelist)

}