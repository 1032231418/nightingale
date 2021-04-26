package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

func main() {
	Host := "https://10.6.100.188:6443"
	BearerToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IlA1anVBaTJlZFppUXVkRXAyeHpabDcxZjZ1QTVPci11dHJoQW4ydVh4NjQifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi10b2tlbi1mcXE3ciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6Ijg0ZjI1ZDg3LTkwOWQtNDljMC1iYmI3LTA3MTQ2Mjg4Mzc1NSIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTphZG1pbiJ9.KwW4Ch_V24_UhUL16cvZ5ZdZKgauZTHixYhx8j9qkdWNEU9dQLrBfHm171UN_cIcbyOdFfXNIpgRLCD7uhW31DSVM1hrBIfkVptoz5KCtjgpOFNEnKxL5iOY0rtqrcXimAPIDZo3KVEVvpy_yCdQ2KKJCnsKyTmnilxmtNrQleiNEZRewW1BxpppySweN01Nk5BKxsV4JuWA3O9cblyNXWXT3YVejk17GOBeWgYinL8qeHJzMWNIQ0id_8oC16UNiBpv31DtZwbdZgYhQDosEP0NyXQD9ewpaA3uQC49tTzZRZGnK6ycYOUPIUjfDhUijWb3W1VzTWp5FpF2zq7cHg"
	config := &rest.Config{
		Host:            Host,
		BearerToken:     BearerToken,
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}
	clientset, err := metrics.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(clientset.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{}))
	podinfos, err := clientset.MetricsV1beta1().PodMetricses("dev").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, podinfo := range podinfos.Items {
		for _, pod := range podinfo.Containers {
			fmt.Println(pod.Name)
			fmt.Println(pod.Usage.Cpu().MilliValue())
			fmt.Println(pod.Usage.Memory().Value() / (1024 * 1024))
		}
	}

}
