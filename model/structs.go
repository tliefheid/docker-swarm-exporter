package model

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/docker/docker/api/types/swarm"
)

type Resources struct {
	NanoCPUs    int64
	MemoryBytes int64
}

func (r Resources) ToNormalCPU() float64 {
	return float64(r.NanoCPUs) / 1e+9
}

// ServiceMetrics contains all service metrics for easy exporting
type ServiceMetrics struct {
	Name        string
	ServiceMode string
	Container   string
	Limits      Resources
	Reservation Resources
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
	Resources     Resources
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

// SwarmMetrics contains generic swarm metrics
type SwarmMetrics struct {
	ID        string
	Container ContainerMetrics
	NCPU      int
	Memory    int64
	Images    int
}

// GetSanitizedID sanitizes the swarm id
func (sm SwarmMetrics) GetSanitizedID() string {
	id := strings.Map(sanitizeRune, sm.ID)
	return id
}

// copied from: https://github.com/census-instrumentation/opencensus-go/blob/master/internal/sanitize.go
// unable to reuse to to restrictions on 'internal' packages
func sanitizeRune(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return r
	}
	// Everything else turns into an underscore
	return '_'
}

// ContainerMetrics holds the counts for serveral container states
type ContainerMetrics struct {
	Running int
	Paused  int
	Stopped int
}
