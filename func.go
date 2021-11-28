package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var ctx = context.Background()

func getData() (GetMonitors, error) {
	cf := getConfig()
	url := "https://api.uptimerobot.com/v2/getMonitors"

	payload := strings.NewReader(fmt.Sprintf("api_key=%v&format=json&logs=1&custom_uptime_ratios=7", cf.Key))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var data GetMonitors
	err := json.Unmarshal(body, &data)

	return data, err
}

func getMonitors(c *fiber.Ctx) error {
	// Initialize Redis
	rdb := initRedis()
	defer rdb.Close()

	byteData, err := rdb.Get(ctx, "status_cache").Bytes()
	var data GetMonitors
	if err != nil {
		log.Println(err)
		data, err = getData()
		if err != nil {
			return err
		}
		byteData, _ = json.Marshal(data)
		rdb.Set(ctx, "status_cache", byteData, 30*time.Second)
	} else {
		err = json.Unmarshal(byteData, &data)
	}

	var ret Ret

	ret.Up = 0
	ret.Total = 0
	ret.Monitors = make([]LightMonitor, 0)

	for _, monitor := range data.Monitors {
		ret.Total += 1
		var m LightMonitor
		m.Url = monitor.Url
		m.Name = monitor.FriendlyName
		m.Ratio = monitor.CustomUptimeRatio
		m.Status = monitor.Status
		if m.Status == 2 {
			ret.Up += 1
		}
		ret.Monitors = append(ret.Monitors, m)
	}

	return c.JSON(ret)
}
