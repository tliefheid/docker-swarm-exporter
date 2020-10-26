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
	fmt.Printf("INFO\n%#v\n", info)

	var sm = model.SwarmMetrics{
		// Container: cm,
		Container: model.ContainerMetrics{
			Running: info.ContainersRunning,
			Paused:  info.ContainersPaused,
			Stopped: info.ContainersStopped,
		},
		NCPU:   info.NCPU,
		Memory: info.MemTotal,
		Images: info.Images,
	}
	common.ExportSwarmClusterInfo(sm)
}
