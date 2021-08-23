package registry

type Registration struct {
	ServiceName string `json:"service_name"`
	ServiceUrl  string `json:"service_url"`
}

func NewRegistration(name, url string) Registration {
	return Registration{
		ServiceName: name,
		ServiceUrl:  url,
	}
}
