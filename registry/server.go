package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	mu            sync.RWMutex
}

type RegService struct {
}

var reg = Registry{
	registrations: make([]Registration, 0, 10),
}

func (r *Registry) add(reg Registration) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.registrations = append(r.registrations, reg)
	//err := r.sendRequiredServices(reg)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (r *Registry) sendRequiredServices(reg Registration) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var p patch
	for _, registration := range r.registrations {
		for _, service := range reg.RequiredServices {
			if registration.ServiceName == service {
				p.added = append(p.added, patchItem{
					name: registration.ServiceName,
					url:  registration.ServiceUrl,
				})
			}
		}
	}

	//err := r.sendPatch(p, reg.ServiceUpdateUrl)
	//return err
	return nil
}

func (r *Registry) sendPatch(p patch, url string) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("send patch info fail with status code: %v", res.StatusCode)
	}

	return nil
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
		err = reg.add(registration)
		if err != nil {
			log.Println("add service failed:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
