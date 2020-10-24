package metrics

// ServiceMetricsStruct contains all service metrics for easy exporting
type ServiceMetrics struct {
	Name           string
	ServiceMode    string
	Container      string
	CPULimit       int64
	MemLimit       int64
	CPUReservation int64
	MemReservation int64
	TimeCreated    int64
	TimeUpdated    int64
	Replicas       float64
}
