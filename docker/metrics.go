package docker

import (
	"github.com/docker/docker/api/types"
	"time"
)

type Metrics struct {
	ContainerId string `json:"cid"`
	Cpu         types.CPUStats
	Memory      types.MemoryStats
	ReadTime    time.Time
	OsType      string `json:"os"`
}

func NewMetrics(c *DockerContainer, stats *types.Stats, osType string) Metrics {
	return Metrics{
		ContainerId: c.ID,
		Cpu:         stats.CPUStats,
		Memory:      stats.MemoryStats,
		ReadTime:    stats.Read,
		OsType:      osType,
	}
}
