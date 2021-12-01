package main

type Config struct {
	Host      string `json:"host"`
	Port      uint   `json:"port"`
	RedisHost string `json:"redis_host"`
	RedisPort uint   `json:"redis_port"`
	Key       string `json:"key"`
}

// JSON return types

type Reason struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

type Log struct {
	Id       uint   `json:"id"`
	Type     uint   `json:"type"`
	Datetime uint   `json:"datetime"`
	Duration int    `json:"duration"`
	Reason   Reason `json:"reason"`
}

type Monitor struct {
	Id                uint   `json:"id"`
	FriendlyName      string `json:"friendly_name"`
	Url               string `json:"url"`
	Type              uint   `json:"type"`
	Interval          uint   `json:"interval"`
	Status            uint   `json:"status"`
	CreateDatetime    uint   `json:"create_datetime"`
	Logs              []Log  `json:"logs"`
	CustomUptimeRatio string `json:"custom_uptime_ratio"`
}

type GetMonitors struct {
	Stat     string    `json:"stat"`
	Monitors []Monitor `json:"monitors"`
}

// Return type

type LightMonitor struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Status uint   `json:"status"`
	Ratio  string `json:"ratio"`
}

type Ret struct {
	Up       uint           `json:"up"`
	Total    uint           `json:"total"`
	Monitors []LightMonitor `json:"monitors"`
}
