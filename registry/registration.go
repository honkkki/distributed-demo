package registry

type Registration struct {
	ServiceName      string `json:"service_name"`
	ServiceUrl       string `json:"service_url"`
	RequiredServices []string
	ServiceUpdateUrl string
}

type patchItem struct {
	name string
	url string
}

type patch struct {
	added []patchItem
	removed []patchItem
}

func NewRegistration(name, url string) Registration {
	return Registration{
		ServiceName: name,
		ServiceUrl:  url,
	}
}
