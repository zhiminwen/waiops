package waiops

type MetricGroup struct {
	Groups []Metric `json:"groups"`
}

type Metric struct {
	Timestamp  int64              `json:"timestamp"`
	ResourceId string             `json:"resourceID"`
	Attributes map[string]string  `json:"attributes"`
	Metrics    map[string]float64 `json:"metrics"`
}
