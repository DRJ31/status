package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DRJ31/status/model"
	"github.com/DRJ31/status/service"
	"io"
	"log"
	"net/http"
)

const (
	WX_MSG_API = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
)

func sendMsg(monitors []model.Monitor) {
	cf := service.GetConfig()
	content := "### UptimeRobot Down Sites"
	for _, monitor := range monitors {
		content += "\n**" + monitor.FriendlyName + "**: " + monitor.Url + "\n"
	}
	loc := WX_MSG_API + cf.WxBotKey

	var msg model.WxMsgMarkdown
	msg.Markdown.Content = content
	msg.Msgtype = "markdown"

	jsonStr, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", loc, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	service.GetConfig()
	gm, err := service.GetData()
	if err != nil {
		panic(err)
	}
	downMonitors := make([]model.Monitor, 0)
	for _, monitor := range gm.Monitors {
		if monitor.Status != 2 {
			downMonitors = append(downMonitors, monitor)
		}
	}
	if len(downMonitors) > 0 {
		sendMsg(downMonitors)
	}
	log.Println("Finish")
}
