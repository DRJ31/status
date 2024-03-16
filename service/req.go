package service

import (
	"encoding/json"
	"fmt"
	"github.com/DRJ31/status/model"
	"net/http"
	"strings"
)

func GetData() (model.GetMonitors, error) {
	url := "https://api.uptimerobot.com/v2/getMonitors"

	payload := strings.NewReader(fmt.Sprintf("api_key=%v&format=json&logs=1&custom_uptime_ratios=7", cfg.Key))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	var data model.GetMonitors
	err := json.NewDecoder(res.Body).Decode(&data)

	return data, err
}
