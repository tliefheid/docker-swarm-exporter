package common

const (
	lblServiceName = "service_name"
	lblContainer   = "container_image"
	lblDeployMode  = "deploy_mode"

	lblState = "state"
	lblID    = "id"

	lblNodeID           = "node_id"
	lblHost             = "hostname"
	lblRole             = "role"
	lblOS               = "os"
	lblArch             = "architecture"
	lblEngVersion       = "engine_version"
	lblNodeState        = "node_state"
	lblManagerReachable = "manager_reachable"
	lblLeader           = "leader"
)

const (
	ServiceModeUnknown    = "unknown"
	ServiceModeReplicated = "replicated"
	ServiceModeGlobal     = "global"
)

const (
	prefixMetric       = ""
	prefixSwarmCluster = "swarm_cluster_"
	prefixSwarmNode    = "swarm_node_"
	prefixServiceSpec  = "swarm_service_spec_"

	info = "info"

	containers = "container_state"
	images     = "images"
	cpu        = "cpu"
	mem        = "memory"

	cpuLimit       = "cpu_limit"
	cpuReservation = "cpu_reservation"
	memLimit       = "memory_limit"
	memReservation = "memory_reservation"

	resourcesCPU = "resources_cpu"
	resourcesMem = "resources_memory"

	timeCreated = "time_created"
	timeUpdated = "time_updated"
	replicas    = "replicas"
)

const (
	swarmInfo              = prefixMetric + prefixSwarmCluster + info
	swarmClusterContainers = prefixMetric + prefixSwarmCluster + containers
	swarmClusterImages     = prefixMetric + prefixSwarmCluster + images
	swarmClusterCPU        = prefixMetric + prefixSwarmCluster + cpu
	swarmClusterMemory     = prefixMetric + prefixSwarmCluster + mem

	swarmNodeInfo           = prefixMetric + prefixSwarmNode + info
	swarmNodeResourceCPU    = prefixMetric + prefixSwarmNode + resourcesCPU
	swarmNodeResourceMemory = prefixMetric + prefixSwarmNode + resourcesMem

	serviceSpecInfo           = prefixMetric + prefixServiceSpec + info
	serviceSpecMemLimit       = prefixMetric + prefixServiceSpec + memLimit
	serviceSpecCPULimit       = prefixMetric + prefixServiceSpec + cpuLimit
	serviceSpecMemReservation = prefixMetric + prefixServiceSpec + memReservation
	serviceSpecCPUReservation = prefixMetric + prefixServiceSpec + cpuReservation
	serviceSpecTimeCreated    = prefixMetric + prefixServiceSpec + timeCreated
	serviceSpecTimeUpdated    = prefixMetric + prefixServiceSpec + timeUpdated
	serviceSpecReplicas       = prefixMetric + prefixServiceSpec + replicas
)
