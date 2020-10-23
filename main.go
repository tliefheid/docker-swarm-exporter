package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// var cli *client.Client

const (
	lblServiceName          = "service_name"
	lblContainer            = "container_image"
	lblDeployMode           = "deploy_mode"
	lblDeployModeReplicated = "replicated"
	lblDeployModeGlobal     = "global"
	lblDeployModeUnknown    = "unknown"
)

var (
	customMetric1 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "foo_total",
		Help: "foobar",
	})
	customMetric2 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bar_total",
		Help: "bar",
	})
	serviceCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "docker_exporter_service_count",
		Help: "number of services in docker",
	})
	serviceInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "service_spec_info",
		Help: "service info",
	}, []string{lblServiceName, lblContainer, lblDeployMode})
	serviceCPULimit = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "service_spec_cpu_limit",
		Help: "service spec cpu limits in CPU Nanos",
	}, []string{lblServiceName})
	serviceCPUReservation = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "service_spec_cpu_reservation",
		Help: "service spec cpu reservations in CPU Nanos",
	}, []string{lblServiceName})
	serviceMemLimit       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "service_spec_memory_limit"}, []string{lblServiceName})
	serviceMemReservation = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "service_spec_memory_reservation"}, []string{lblServiceName})
	serviceTimeUpdate     = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "service_spec_time_last_update"}, []string{lblServiceName})
	serviceTimeCreated    = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "service_spec_time_created"}, []string{lblServiceName})

	swarmNodeInfo            = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "swarm_node_info"}, []string{"nodeid", "role", "hostname", "os", "architecture", "availability", "engine_version", "nodestate", "manager", "leader", "manager_state"})
	swarmNodeResourcesMemory = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "swarm_node_resources_memory"}, []string{"nodeid"})
	swarmNodeResourcesCPU    = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "swarm_node_resources_cpu"}, []string{"nodeid"})

	swarmInfo            = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "swarm_info"}, []string{"os", "kernel"})
	swarmContainers      = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "swarm_containers"}, []string{"state"})
	swarmResourcesCPU    = promauto.NewGauge(prometheus.GaugeOpts{Name: "swarm_resources_cpu"})
	swarmResourcesMemory = promauto.NewGauge(prometheus.GaugeOpts{Name: "swarm_resources_memory"})
)

func collectMetrics() {
	collectDockerServices()
	collectDockerNode()
	collectDockerSwarm()
	collectDockerExtra()
	customMetric2.Inc()
}

func collectDockerSwarm() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	info, err := cli.Info(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nINFO\n%#v\n\n", info)

	containersRunning := info.ContainersRunning
	containersPaused := info.ContainersPaused
	containersStopped := info.ContainersStopped
	swarmContainers.WithLabelValues("running").Set(float64(containersRunning))
	swarmContainers.WithLabelValues("paused").Set(float64(containersPaused))
	swarmContainers.WithLabelValues("stopped").Set(float64(containersStopped))
	swarmResourcesCPU.Set(float64(info.NCPU))
	swarmResourcesMemory.Set(float64(info.MemTotal))
}

func collectDockerNode() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}

	for _, n := range nodes {
		fmt.Printf("NODE:\n%#v\n", n)
		nodeid := n.ID
		role := string(n.Spec.Role)
		hostname := n.Description.Hostname
		os := n.Description.Platform.OS
		architecture := n.Description.Platform.Architecture
		availability := string(n.Spec.Availability)
		engineVersion := n.Description.Engine.EngineVersion
		nodeState := string(n.Status.State)
		isManager := (n.ManagerStatus != nil)
		// managerLeader := "false"
		manager := "false"
		managerReachable := "unknown"
		managerLeader := "false"
		if isManager {
			manager = "true"
			managerReachable = string(n.ManagerStatus.Reachability)
			managerLeader = strconv.FormatBool(n.ManagerStatus.Leader)
		}
		cpu := n.Description.Resources.NanoCPUs
		memory := n.Description.Resources.MemoryBytes

		swarmNodeInfo.WithLabelValues(nodeid, role, hostname, os, architecture, availability, engineVersion, nodeState, manager, managerLeader, managerReachable).Set(1)
		swarmNodeResourcesCPU.WithLabelValues(nodeid).Set(float64(toNormalCPU(cpu)))
		swarmNodeResourcesMemory.WithLabelValues(nodeid).Set(float64(memory))
	}

}
func collectDockerExtra() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// node, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	// fmt.Printf("NODES\n%#v\n\n", node)

	swarm, err := cli.SwarmInspect(context.Background())
	fmt.Printf("\nSWARM\n%#v\n\n", swarm)
}
func collectDockerServices() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	_serviceCount := len(services)
	fmt.Printf("found '%d' services\n", _serviceCount)
	serviceCount.Set(float64(_serviceCount))

	for _, s := range services {
		spec := s.Spec
		serviceMode := spec.Mode
		task := spec.TaskTemplate
		resources := task.Resources
		limits := resources.Limits
		reservation := resources.Reservations
		serviceName := s.Spec.Annotations.Name
		container := task.ContainerSpec.Image

		created := s.Meta.CreatedAt
		lastUpdate := s.Meta.UpdatedAt
		fmt.Println("\n----------------------------------------")
		fmt.Printf("Service: %s\n", s.Spec.Annotations.Name)
		fmt.Printf("container: %s\n", container)
		fmt.Println("----------------------------------------")
		// fmt.Printf("Mode: %+v\n", serviceMode)
		// fmt.Printf("replicated: %+v\n", serviceMode.Replicated)
		// if serviceMode.Replicated != nil {
		// 	rep := *serviceMode.Replicated.Replicas
		// 	fmt.Printf("replicas: %+v\n", serviceMode.Replicated.Replicas)
		// 	fmt.Printf("replicas: %d\n", *serviceMode.Replicated.Replicas)
		// 	fmt.Printf("replicas: %d\n", rep)
		// 	fmt.Printf("replicas: %s\n", strconv.FormatUint(rep, 10))
		// }
		// fmt.Printf("global: %+v\n", serviceMode.Global)
		// fmt.Println("----------------------------------------")
		// fmt.Printf("TaskTemplate: %+v\n", spec.TaskTemplate)
		// fmt.Printf("Resources: %+v\n", resources)
		// fmt.Printf("Limits: %+v\n", limits)
		// fmt.Printf("Reservations: %+v\n", reservation)
		// fmt.Println("----------------------------------------")
		// fmt.Printf("T1        : %+v\n", created)
		// fmt.Printf("t2        : %+v\n", created.Unix())
		// fmt.Printf("NANO CPU         : %+v\n", limits.NanoCPUs)
		// fmt.Printf("NANO CPU normal  : %+v\n", toNormalCPU(limits.NanoCPUs))
		fmt.Println("----------------------------------------")
		fmt.Printf("%#v\n", s)
		fmt.Println("----------------------------------------")
		// lbl := prometheus.Labels{"name": serviceName, "container": container}
		// serviceInfo.With(lbl).Set(1)

		serviceCPULimit.WithLabelValues(serviceName).Set(float64(toNormalCPU(limits.NanoCPUs)))
		serviceCPUReservation.WithLabelValues(serviceName).Set(float64(toNormalCPU(reservation.NanoCPUs)))

		serviceMemLimit.WithLabelValues(serviceName).Set(float64(limits.MemoryBytes))
		serviceMemReservation.WithLabelValues(serviceName).Set(float64(reservation.MemoryBytes))

		serviceTimeCreated.WithLabelValues(serviceName).Set(float64(created.Unix()))
		serviceTimeUpdate.WithLabelValues(serviceName).Set(float64(lastUpdate.Unix()))

		serviceModeLbl := lblDeployModeUnknown
		if serviceMode.Global != nil {
			serviceModeLbl = lblDeployModeGlobal
		} else if serviceMode.Replicated != nil {
			serviceModeLbl = lblDeployModeReplicated
			// set separate gauge to number of replicas available
			//).Set(float64(*serviceMode.Replicated.Replicas))
		}
		serviceInfo.WithLabelValues(serviceName, container, serviceModeLbl).Set(1)
		// cli.Close()

	}

}
func toNormalCPU(nanoCPU int64) float64 {
	return float64(nanoCPU) / 1e+9
}

func httpWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		collectMetrics()
		h.ServeHTTP(w, r)
	})
}
func cleanup() {
	fmt.Println("cleanup")
	// cli.Close()
}
func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	// var err error
	// cli, err = client.NewEnvClient()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("===============================================")
	collectDockerServices()
	fmt.Println("===============================================")
	collectDockerNode()
	fmt.Println("===============================================")
	collectDockerSwarm()
	fmt.Println("===============================================")
	collectDockerExtra()
	fmt.Println("===============================================")

	// customMetric1.Inc()
	// http.Handle("/metrics", promhttp.Handler())
	http.Handle("/metrics", httpWrapper(promhttp.Handler()))
	http.ListenAndServe(":2112", nil)

}
