package services

import (
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func createPodInfo(addresses []v1.EndpointAddress) *podInfo {
	names := make([]string, 0)
	for _, pod := range addresses {
		names = append(names, pod.TargetRef.Name)
	}

	return &podInfo{
		Total: len(addresses),
		Pods:  names,
	}
}


func GetServiceStatus(client corev1.CoreV1Interface, namespace, name string) (*componentEndpoints, error) {
	endpoint, err := client.Endpoints(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	subset := endpoint.Subsets[0]

	notReady := subset.NotReadyAddresses
	notReadyPods := createPodInfo(notReady)

	ready := subset.Addresses
	readyPods := createPodInfo(ready)

	resp := componentEndpoints{
		Total:        len(ready) + len(notReady),
		Available:    readyPods,
		NotAvailable: notReadyPods,
	}

	return &resp, nil
}
