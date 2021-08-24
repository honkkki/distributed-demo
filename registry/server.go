package registry

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const (
	ServerPort = ":9099"
	ServerUrl  = "http://127.0.0.1" + ServerPort + "/services"
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

func (r *Registry) remove(url string) error {
	for i, registration := range r.registrations {
		if registration.ServiceUrl == url {
			r.mu.Lock()
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			r.mu.Unlock()
			return nil
		}
	}

	return nil
}

func (r *RegService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
	case http.MethodDelete:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(body)
		log.Printf("Removing service url: %s\n", url)
		err = reg.remove(url)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println("remove service success!")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
