package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	api "github.com/docker/docker/client"
)

type MetricsCollector interface {
	Start()
	Stop()
}

type DockerContainer struct {
	*api.Client
	types.Container
	metricsChan chan Metrics // channel used to pass metrics to logger
}

func dockerContainer(c types.Container) *DockerContainer {
	return &DockerContainer{
		Container: c,
	}
}

func (c *DockerContainer) SetMetricsChan(channel chan Metrics) {
	c.metricsChan = channel
}

func (c *DockerContainer) Start() {
	//@todo: Error handling
	go func() {
		var buf bytes.Buffer
		for {
			//@todo: wtf does stream option do???
			stats, _ := c.ContainerStats(context.Background(), c.ID, false)
			buf.ReadFrom(stats.Body)
			var cStats types.Stats
			//@todo: Error handling
			json.Unmarshal(buf.Bytes(), &cStats)
			c.metricsChan <- NewMetrics(c, &cStats, stats.OSType)
			buf.Reset()
		}
	}()
}

func GetAllRunningContainers() []*DockerContainer {
	cli, err := api.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	dcList := make([]*DockerContainer, 0, len(containers))
	for _, c := range containers {
		dc := dockerContainer(c)
		dc.Client = cli
		dcList = append(dcList, dc)
	}
	return dcList
}
