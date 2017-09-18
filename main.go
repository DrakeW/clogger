package main

import (
	"encoding/json"
	"github.com/DrakeW/clogger/docker"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	containers := docker.GetAllRunningContainers()
	metricsChan := make(chan docker.Metrics)
	for _, c := range containers {
		c.SetMetricsChan(metricsChan)
		c.Start()
	}
	for metrics := range metricsChan {
		content, _ := json.Marshal(metrics)
		sugar.Infof("[Metrics] %s", string(content))
	}
}
