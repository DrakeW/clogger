package docker

import (
	"github.com/docker/docker/api/types"
	"time"
)

type Metrics struct {
	cpu      types.CPUStats
	memory   types.MemoryStats
	readTime time.Time
	osType   string
}

func NewMetrics(stats *types.Stats, osType string) Metrics {
	return Metrics{
		cpu:      stats.CPUStats,
		memory:   stats.MemoryStats,
		readTime: stats.Read,
		osType:   osType,
	}
}
