package main

import (
	"github.com/DrakeW/clogger/docker"
	"github.com/DrakeW/clogger/logging"
)

func main() {
	containers := docker.GetAllRunningContainers()
	metricsChan := make(chan docker.Metrics)
	for _, c := range containers {
		c.SetMetricsChan(metricsChan)
		c.Start()
	}
	defaultTransformer := logging.NewDefaultTransformer("default_transformer_1")
	defaultTransformer.StartTransform(metricsChan)
	defaultTransformer.Output()
}
