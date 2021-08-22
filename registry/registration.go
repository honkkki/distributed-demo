package registry

type Registration struct {
	ServiceName string `json:"service_name"`
	ServiceUrl  string `json:"service_url"`
}

const (
	LogService = "LogService"
)
