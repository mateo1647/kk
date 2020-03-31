package resources

import (
	"fmt"
	"strings"

	"github.com/mateo1647/kk/internal/options"
	"github.com/mateo1647/kk/pkg/client"

	"github.com/mateo1647/kk/util"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	clientset = client.InitClient()
)

// Services - a public function for searching services with keyword
func GetServicesandPods(opt *options.SearchOptions, keyword string) []GetServicesandPodsResponse {
	//ns, o := util.SetOptions(opt)
	var serviceResponse []GetServicesandPodsResponse
	serviceList := util.ServiceList(opt)
	for _, service := range serviceList.Items {
		selector := service.Spec.Selector
		if len(keyword) > 0 {
			match := strings.Contains(service.Name, keyword)
			if !match {
				continue
			}
		}
		if len(selector) > 0 {
			ns, _ := util.SetOptions(opt)
			podList, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{LabelSelector: util.KeysString(selector)})
			if err != nil {
				log.WithFields(log.Fields{
					"err": err.Error(),
				}).Debug("Unable to get service and pod List")
			}
			var podResponse []PodResponse
			for _, pod := range podList.Items {
				podResponse = append(podResponse, NewPodDetails(pod))

			}
			headerLine := fmt.Sprintf(util.ServiceHeader)
			serviceInfo := GetServicesandPodsResponse{Service: service, Headerline: headerLine, PodResponse: podResponse}
			serviceResponse = append(serviceResponse, serviceInfo)
		}
	}
	return serviceResponse
}

type PodResponse struct {
	HeaderLine string

	StatusLine string
}

func NewPodDetails(pod v1.Pod) PodResponse {
	total := len(pod.Status.ContainerStatuses)
	var ready int
	var restarts int32
	for _, c := range pod.Status.ContainerStatuses {
		if c.Ready {
			ready++
		}
		restarts += c.RestartCount
	}
	//var statusLine string
	//buf := bytes.NewBuffer(nil)
	age := &util.Age{Time: pod.Status.StartTime.Time}
	//w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)
	//headerLine := fmt.Sprintf(util.ServiceHeader)
	statusLine := fmt.Sprintf(util.ServiceRowTemplate,
		pod.Namespace,
		pod.Name,
		ready,
		total,
		pod.Status.Phase,
		restarts,
		age.Relative())

	return PodResponse{StatusLine: statusLine}
}

type GetServicesandPodsResponse struct {
	Service     v1.Service
	Headerline  string
	PodResponse []PodResponse
}
