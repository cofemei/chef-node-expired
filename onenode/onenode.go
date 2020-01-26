package main

import (
	// "encoding/json"
	"fmt"
	"os"
	"github.com/go-chef/chef"
	"io/ioutil"
	// "encoding/json"
)

type ChefPackages struct {
	Chef struct {
		ChefRoot string `json:"chef_root"`
		Version string `json:"version"`
	}  `json:"chef"`
	Ohai struct {
		Ohai string `json:"ohai"`
		Version string `json:"version"`
	}  `json:"ohai"`
}

func main() {
	ef, err := ioutil.ReadFile("george.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	client, err := chef.NewClient(&chef.Config{
		Name:    "george",
		Key:     string(ef),
		BaseURL: "https://chef.com.tw/organizations/se/",
	})

	serverNode, _ := client.Nodes.Get("stg-ecfe-spot-i-09f6caf353781d1ba")

	myMap := serverNode.AutomaticAttributes["cloud"].(map[string]interface{})
	fmt.Println(myMap)
	fmt.Println(myMap["local_ipv4"])

	// myMap2 := myMap["local_ipv4"].(map[string]interface{})
	// fmt.Println(myMap2)

	// var v ChefPackages
	// inrec, _ := json.Marshal( serverNode.AutomaticAttributes["chef_packages"] )
	// fmt.Println(inrec)
	// json.Unmarshal(inrec, &v)
	// fmt.Println(v.Chef.Version)

	// jsonData, err := json.MarshalIndent(serverNode, "", "\t")
	// fmt.Println(err)
	// os.Stdout.Write(jsonData)
	// os.Stdout.WriteString("\n")

}
