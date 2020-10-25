package model

import (
	"strconv"

	"github.com/docker/docker/api/types/swarm"
)

type resources struct {
	NanoCPUs    int64
	MemoryBytes int64
}

func (r resources) ToNormalCPU() float64 {
	return float64(r.NanoCPUs) / 1e+9
}

// ServiceMetrics contains all service metrics for easy exporting
type ServiceMetrics struct {
	Name        string
	ServiceMode string
	Container   string
	Limits      resources
	Reservation resources
	TimeCreated int64
	TimeUpdated int64
	Replicas    float64
}

// NodeMetrics contains all node metrics for easy exporting
type NodeMetrics struct {
	ID            string
	Host          string
	Role          swarm.NodeRole
	Os            string
	Architecture  string
	Availability  string
	EngineVersion string
	NodeStatus    string
	Resources     resources
	ManagerInfo   managerInfo
}

type managerInfo struct {
	Reachability string
	Leader       bool
}

// ManagerReachability get the manager reachability if ManagerInfo != nil
func (nm NodeMetrics) ManagerReachability() string {
	if nm.ManagerInfo == (managerInfo{}) {
		return string(swarm.ReachabilityUnknown)
	}
	return string(nm.ManagerInfo.Reachability)
}

// IsLeader get the bool if a node is a manager, if ManagerInfo != nil
func (nm NodeMetrics) IsLeader() string {
	if nm.ManagerInfo == (managerInfo{}) {
		return "false"
	}
	return strconv.FormatBool(nm.ManagerInfo.Leader)
}

type SwarmMetrics struct {
	ID        string
	Container ContainerMetrics
	NCPU      int
	Memory    int64
	Images    int
}
type ContainerMetrics struct {
	Running int
	Paused  int
	Stopped int
}
