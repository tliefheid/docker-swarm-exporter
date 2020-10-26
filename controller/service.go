package controller

import (
	"fmt"

	"github.com/toml-dev/docker-swarm-exporter/common"
	"github.com/toml-dev/docker-swarm-exporter/model"
)

// UpdateServiceMetrics gathers service metrics and expose them to prometheus
func UpdateServiceMetrics() {
	fmt.Println("Update Service Metrics")

	services := common.GetServiceList()

	var serviceArr []model.ServiceMetrics
	// _serviceCount := len(services)
	// fmt.Printf("\nfound '%d' services\n", _serviceCount)
	for _, s := range services {
		// serviceName := s.Spec.Annotations.Name
		// fmt.Printf("%s", serviceName)

		actualServiceMode := s.Spec.Mode
		replicas := float64(-1)
		serviceMode := common.ServiceModeUnknown
		if actualServiceMode.Global != nil {
			serviceMode = common.ServiceModeGlobal
		} else if actualServiceMode.Replicated != nil {
			serviceMode = common.ServiceModeReplicated
			replicas = float64(*s.Spec.Mode.Replicated.Replicas)
		}
		var sm = model.ServiceMetrics{
			Name:        s.Spec.Annotations.Name,
			ServiceMode: serviceMode,
			Container:   s.Spec.TaskTemplate.ContainerSpec.Image,
			TimeCreated: s.Meta.CreatedAt.Unix(),
			TimeUpdated: s.Meta.UpdatedAt.Unix(),
			Replicas:    replicas,
		}
		sm.Limits.NanoCPUs = s.Spec.TaskTemplate.Resources.Limits.NanoCPUs
		sm.Limits.MemoryBytes = s.Spec.TaskTemplate.Resources.Limits.MemoryBytes
		sm.Reservation.NanoCPUs = s.Spec.TaskTemplate.Resources.Reservations.NanoCPUs
		sm.Reservation.MemoryBytes = s.Spec.TaskTemplate.Resources.Reservations.MemoryBytes

		// fmt.Printf("%#v\n", sm)

		serviceArr = append(serviceArr, sm)
	}

	common.ExportServiceMetrics(serviceArr)
}
