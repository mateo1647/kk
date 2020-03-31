package resources

import (
	"strings"

	"github.com/mateo1647/kk/internal/options"
	"github.com/mateo1647/kk/util"
	v1 "k8s.io/api/core/v1"
)

// Services - a public function for searching services with keyword
func GetServices(opt *options.SearchOptions, keyword string) []GetServicesResponse {
	var serviceResponse []GetServicesResponse
	serviceList := util.ServiceList(opt)

	for _, service := range serviceList.Items {
		// return all services under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(service.Name, keyword)
			if !match {
				continue
			}
		}
		serviceInfo := GetServicesResponse{
			Service: service,
		}
		serviceResponse = append(serviceResponse, serviceInfo)
	}
	return serviceResponse
}

type GetServicesResponse struct {
	Service v1.Service
}
