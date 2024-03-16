package main

import (
	"context"
	"encoding/json"
	"github.com/DRJ31/status/model"
	"github.com/DRJ31/status/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

var ctx = context.Background()

func getMonitors(c *fiber.Ctx) error {
	// Initialize Redis
	rdb := service.InitRedis()
	defer rdb.Close()

	byteData, err := rdb.Get(ctx, "status_cache").Bytes()
	var data model.GetMonitors
	if err != nil {
		log.Println(err)
		data, err = service.GetData()
		if err != nil {
			return err
		}
		byteData, _ = json.Marshal(data)
		rdb.Set(ctx, "status_cache", byteData, 30*time.Second)
	} else {
		err = json.Unmarshal(byteData, &data)
	}

	var ret model.Ret

	ret.Up = 0
	ret.Total = 0
	ret.Monitors = make([]model.LightMonitor, 0)

	for _, monitor := range data.Monitors {
		ret.Total += 1
		var m model.LightMonitor
		m.Url = monitor.Url
		m.Name = monitor.FriendlyName
		m.Ratio = monitor.CustomUptimeRatio
		m.Status = monitor.Status
		m.Logs = monitor.Logs
		if m.Status == 2 {
			ret.Up += 1
		}
		ret.Monitors = append(ret.Monitors, m)
	}

	return c.JSON(ret)
}
