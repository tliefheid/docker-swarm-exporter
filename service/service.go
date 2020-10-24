package service

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/toml-dev/docker-swarm-exporter/common"
	"github.com/toml-dev/docker-swarm-exporter/metrics"
)

func UpdateServiceMetrics() {
	fmt.Println("Update Service Metrics")
	// get cli √
	// collect metrics √
	// fill struct √
	// send struct to export
	cli := common.GetCLI()
	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	var serviceArr []metrics.ServiceMetrics
	printSlice(serviceArr)
	_serviceCount := len(services)
	fmt.Printf("found '%d' services\n", _serviceCount)
	for _, s := range services {
		// sm = metrics.ServiceMetrics {
		// 	name: s.Spec.Annotations.Name
		// }
		// serviceArr.append(sm)
		serviceName := s.Spec.Annotations.Name
		fmt.Printf("%s", serviceName)

		actualServiceMode := s.Spec.Mode
		replicas := float64(-1)
		serviceMode := common.ServiceModeUnknown
		if actualServiceMode.Global != nil {
			serviceMode = common.ServiceModeGlobal
		} else if actualServiceMode.Replicated != nil {
			serviceMode = common.ServiceModeReplicated
			replicas = float64(*s.Spec.Mode.Replicated.Replicas)
		}
		var sm = metrics.ServiceMetricsStruct{
			Name:           s.Spec.Annotations.Name,
			ServiceMode:    serviceMode,
			Container:      s.Spec.TaskTemplate.ContainerSpec.Image,
			CPULimit:       s.Spec.TaskTemplate.Resources.Limits.NanoCPUs,
			CPUReservation: s.Spec.TaskTemplate.Resources.Reservations.NanoCPUs,
			MemLimit:       s.Spec.TaskTemplate.Resources.Limits.MemoryBytes,
			MemReservation: s.Spec.TaskTemplate.Resources.Reservations.MemoryBytes,
			TimeCreated:    s.Meta.CreatedAt.Unix(),
			TimeUpdated:    s.Meta.UpdatedAt.Unix(),
			Replicas:       replicas,
		}
		fmt.Printf("%#v\n", sm)

		serviceArr = append(serviceArr, sm)
	}
	printSlice(serviceArr)

	cli.Close()

	common.ExportServiceMetrics(serviceArr)
}

func printSlice(s []metrics.ServiceMetricsStruct) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
