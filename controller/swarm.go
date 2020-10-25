package controller

import (
	"fmt"

	"github.com/toml-dev/docker-swarm-exporter/common"
	"github.com/toml-dev/docker-swarm-exporter/model"
)

// UpdateSwarmMetrics gathers swarm metrics and expose them to prometheus
func UpdateSwarmMetrics() {
	fmt.Println("Update Swarm Info Metrics")
	info := common.GetSwarmInfo()

	var cm = model.ContainerMetrics{}

	var sm = model.SwarmMetrics{
		// Container: cm,
		Container: model.ContainerMetrics{
			Running: info.ContainersRunning,
			Paused:  info.ContainersPaused,
			Stopped: info.ContainersStopped,
		},
	}
	sm.Container.Paused = info.ContainersPaused
	// sm.Container
}
