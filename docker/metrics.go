package docker

import (
	"github.com/docker/docker/api/types"
	"time"
)

type Metrics struct {
	containerId string `json:"cid"`
	cpu         types.CPUStats
	memory      types.MemoryStats
	readTime    time.Time
	osType      string `json:"os"`
}

func NewMetrics(c *DockerContainer, stats *types.Stats, osType string) Metrics {
	return Metrics{
		containerId: c.ID,
		cpu:         stats.CPUStats,
		memory:      stats.MemoryStats,
		readTime:    stats.Read,
		osType:      osType,
	}
}
