package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DRJ31/status/model"
	"github.com/DRJ31/status/service"
	"github.com/go-redis/redis/v8"
	"io"
	"log"
	"net/http"
	"time"
)

var ctx context.Context

const (
	WX_MSG_API = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
	REDIS_KEY  = "uptime_robot_record"
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

func handleDownSites(rdb *redis.Client, monitors []model.Monitor) {
	if len(monitors) > 0 {
		monitorsByte, _ := json.Marshal(monitors)
		// Check if message has sent within an hour
		cachedMonitors, err := rdb.Get(ctx, REDIS_KEY).Bytes()
		if errors.Is(err, redis.Nil) {
			sendMsg(monitors)
			rdb.Set(ctx, REDIS_KEY, monitorsByte, time.Hour)
			return
		}

		// Check the content of monitor in redis
		var cms []model.Monitor
		var incomplete bool
		err = json.Unmarshal(cachedMonitors, &cms)
		if err != nil {
			sendMsg(monitors)
			rdb.Set(ctx, REDIS_KEY, monitorsByte, time.Hour)
			return
		}

		// Check if the two lists are the same
		for _, cm := range cms {
			incomplete = true
			for _, m := range monitors {
				if cm.Url == m.Url {
					incomplete = false
					break
				}
			}
			if incomplete {
				sendMsg(monitors)
				rdb.Set(ctx, REDIS_KEY, monitorsByte, time.Hour)
				return
			}
		}

	}
}

func main() {
	rdb := service.InitRedis()
	defer rdb.Close()

	// Get data
	gm, err := service.GetData()
	if err != nil {
		panic(err)
	}

	// Filter down websites
	downMonitors := make([]model.Monitor, 0)
	for _, monitor := range gm.Monitors {
		if monitor.Status != 2 {
			downMonitors = append(downMonitors, monitor)
		}
	}

	handleDownSites(rdb, downMonitors)
	log.Println("Finish")
}
