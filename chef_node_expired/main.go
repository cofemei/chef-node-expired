package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/go-chef/chef"
	"github.com/jedib0t/go-pretty/table"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"encoding/json"
	"os"
	"sort"
	"time"
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

func postSlack(text string, slackHooksURL string, channel string) {
	sl := NewSlack(slackHooksURL, fmt.Sprintf("```%s```", text), "webhookbot", ":eyes:", "", channel)
	sl.Send()
}

func postRawSlack(text string, slackHooksURL string, channel string) {
	sl := NewSlack(slackHooksURL, fmt.Sprintf("%s", text), "webhookbot", ":eyes:", "", channel)
	sl.Send()
}

// ohaiInt64, expiredSec, isExpired, chef version, ipv4, error
func isExpired(node chef.Node, threshold int) (int64, int, bool, string, string, bool) {
	if node.AutomaticAttributes["ohai_time"] == nil {
		return 0, 0, true, "", "", false
	}
	// chef-client version
	var v ChefPackages
	inrec, _ := json.Marshal( node.AutomaticAttributes["chef_packages"] )
	json.Unmarshal(inrec, &v)
	// ip
	var ipv4 string
	
	if node.AutomaticAttributes["cloud"] != nil {
		cloudIPMap := node.AutomaticAttributes["cloud"].(map[string]interface{})
		ipv4 = cloudIPMap["local_ipv4"].(string)
	} else {
		ipv4 = ""
	}
	
	ohaiInt64, ok := node.AutomaticAttributes["ohai_time"].(float64)
	if ok {
		return int64(ohaiInt64), (int)(time.Now().Unix() - int64(ohaiInt64)), (int)(time.Now().Unix()-int64(ohaiInt64)) > (threshold * 60 * 60), v.Chef.Version, ipv4, false
	}
	return 0, 0, false, "", "", true
}

func nodeCheck(client *chef.Client, nodeName string, url string, c chan ExpiredNode) {
	serverNode, _ := client.Nodes.Get(nodeName)
	ohaiTime, expiredSec, isExpired, v, ipv4, _ := isExpired(serverNode, 6)
	var os string
	if serverNode.AutomaticAttributes["os"] == nil {
		os = "unknow"
	} else {
		os = serverNode.AutomaticAttributes["os"].(string)
	}
	c <- ExpiredNode{Version: v, ExpiredSec: expiredSec, NodeName: nodeName, IsExpired: isExpired, OS: os, URL: url, OhaiTime: ohaiTime, ipv4: ipv4}
}



func output(elist ExpiredNodeList, slackHooksURL string, channel string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	postRawSlack(":bangbang: Begin Chef client 未執行 Node List", slackHooksURL, channel )
	t.AppendHeader(table.Row{"#", "EXPRIED Check In", "Node", "OS", "URL", "Version", "ip"})
	resultNum := 0 // all linux resout host
	for i := range elist {
		node := elist[i]
		if node.IsExpired && (node.OS == "linux" || node.OS == "unknow") {
			resultNum++
			t.AppendRows([]table.Row{
				{resultNum, humanize.Time(time.Unix(node.OhaiTime, 0)), node.NodeName, node.OS, node.URL, node.Version, node.ipv4},
			})
			if (resultNum % 5) == 0 {
				postSlack(t.Render(), slackHooksURL, channel )
				t.Render()
				t = table.NewWriter()
				t.SetOutputMirror(os.Stdout)
			}
		}
	}

	if t.Length() != 0 {
		t.Render()
		postSlack(t.Render(), slackHooksURL, channel )
	}
	postRawSlack(":bangbang: Finished Chef client 未執行 Node List", slackHooksURL, channel )

}

func outputAll(elist ExpiredNodeList, slackHooksURL string, channel string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	postRawSlack(":bangbang: Chef client checkIn Node List", slackHooksURL, channel )
	t.AppendHeader(table.Row{"#", "EXPRIED Check In", "Node", "OS", "URL", "Version", "ip"})
	resultNum := 0 // all linux resout host
	for i := range elist {
		node := elist[i]
		resultNum++
		t.AppendRows([]table.Row{
			{resultNum, humanize.Time(time.Unix(node.OhaiTime, 0)), node.NodeName, node.OS, node.URL, node.Version, node.ipv4},
		})
	}

	if t.Length() != 0 {
		t.Render()
		postSlack(t.Render(), slackHooksURL, channel )
	}
}

func main() {
	_ = godotenv.Load("env.sh")
	filename := "encrypted_pem.txt"
	var profile, region string
	if os.Getenv("PROFILE") != "" {
		profile = os.Getenv("PROFILE")
	}
	fmt.Println(os.Getenv("PROFILE"))
	if os.Getenv("REGION") != "" {
		region = os.Getenv("REGION")
	}
	key, err := getpem(profile, region, filename)
	if err != nil {
		fmt.Println("Issue encrypted pem Decrypt:", err)
	}
	client, err := chef.NewClient(&chef.Config{
		Name:    os.Getenv("USERNAME"),
		Key:     key,
		BaseURL: os.Getenv("CHEF_SERVER_URL"),
	})
	if err != nil {
		fmt.Println("Issue setting up client:", err)
	}
	nodemap, err := client.Nodes.List()
	if err != nil {
		fmt.Println("Issue listing Nodes:", err)
	}
	ch := make(chan ExpiredNode)
	for k, v := range nodemap {
		go nodeCheck(client, k, v, ch)
	}
	var elist ExpiredNodeList
	for i := 0; i < len(nodemap); i++ {
		elist = append(elist, <-ch)
	}
	sort.Sort(elist)
	output(elist, os.Getenv("SLACK_HOOKS_URL"), os.Getenv("CHANNEL"))

	lambda.Start(LambdaHandler)
}

func LambdaHandler() (string, error) {
	return fmt.Sprintf("Hello !", ), nil
}