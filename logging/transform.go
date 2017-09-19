package logging

import (
	"encoding/json"
	"github.com/DrakeW/clogger/docker"
	"github.com/docker/docker/api/types"
	"go.uber.org/zap"
)

type Transformer interface {
	StartTransform(chan docker.Metrics)
	Output() chan docker.Metrics // Output Chan
}

type DefaultTransformer struct {
	name       string
	logger     *zap.SugaredLogger
	outputChan chan []byte
}

func NewDefaultTransformer(name string) *DefaultTransformer {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return &DefaultTransformer{
		name:       name,
		logger:     sugar,
		outputChan: make(chan []byte),
	}
}

func (t *DefaultTransformer) StartTransform(inputChan chan docker.Metrics) {
	go func() {
		for metric := range inputChan {
			mJson, _ := json.Marshal(struct {
				containerId string `json:"container_id"`
				osType      string `json:"os_type"`
				*types.CPUStats
				*types.MemoryStats
			}{
				metric.ContainerId,
				metric.OsType,
				&metric.Cpu,
				&metric.Memory,
			})
			t.outputChan <- mJson
		}
	}()
}

func (t *DefaultTransformer) Output() {
	for mJson := range t.outputChan {
		t.logger.Infof("%s", string(mJson))
	}
}
