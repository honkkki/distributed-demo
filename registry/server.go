package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const (
	ServerPort = ":9099"
	ServerUrl  = "127.0.0.1" + ServerPort + "/services"
)

type Registry struct {
	registrations []Registration
	mu            sync.Mutex
}

type RegService struct {
}

var reg = Registry{
	registrations: make([]Registration, 0, 10),
}

func (r *Registry) add(reg Registration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.registrations = append(r.registrations, reg)
}

func (r *RegService) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	log.Println("request received")
	switch req.Method {
	case http.MethodPost:
		dec := json.NewDecoder(req.Body)
		var registration Registration
		err := dec.Decode(&registration)
		if err != nil {
			log.Println("json decode error:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: [%s] with URL: %s\n", registration.ServiceName, registration.ServiceUrl)
		reg.add(registration)
		log.Println("add service success!")
		log.Println(reg.registrations)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
