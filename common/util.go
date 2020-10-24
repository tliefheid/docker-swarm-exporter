package common

func toNormalCPU(nanoCPU int64) float64 {
	return float64(nanoCPU) / 1e+9
}