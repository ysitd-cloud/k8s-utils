package services // import "code.ysitd.cloud/k8s/utils/go/services"

type componentEndpoints struct {
	Total        int      `json:"total"`
	Available    *podInfo `json:"available"`
	NotAvailable *podInfo `json:"not_available"`
}

type podInfo struct {
	Total int      `json:"total"`
	Pods  []string `json:"pods"`
}
