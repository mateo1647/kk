package resources

import (
	"strings"

	"github.com/mateo1647/kk/internal/options"
	"github.com/mateo1647/kk/util"
	corev1 "k8s.io/api/core/v1"
)

func GetPods(opt *options.SearchOptions, keyword string) []GetPodsResponse {
	var podResponse []GetPodsResponse
	podList := util.PodList(opt)

	for _, pod := range podList.Items {
		// return all services under namespace if no keyword specific
		if len(keyword) > 0 {
			match := strings.Contains(pod.Name, keyword)
			if !match {
				continue
			}
		}
		podInfo := GetPodsResponse{
			Pod: pod,
		}
		podResponse = append(podResponse, podInfo)
	}
	return podResponse
}

type GetPodsResponse struct {
	Pod corev1.Pod
}
