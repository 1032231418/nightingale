// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcs

import (
	"context"
	"fmt"
	"github.com/didi/nightingale/v4/src/common/dataobj"
	config2 "github.com/didi/nightingale/v4/src/modules/agentd/config"
	"github.com/didi/nightingale/v4/src/modules/agentd/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

func PodMetrics() []*dataobj.MetricValue {
	Host := config2.Config.K8sClient.Apihost
	BearerToken := config2.Config.K8sClient.Apitoken
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

	var ret []*dataobj.MetricValue

	for _, podinfo := range podinfos.Items {
		for _, pod := range podinfo.Containers {
			tags := fmt.Sprintf("pod_name=%s", pod.Name)
			ret = append(ret, core.GaugeValue("pod.mem.used", pod.Usage.Memory().Value()/(1024*1024), tags))
			ret = append(ret, core.GaugeValue("pod.cpu.used", pod.Usage.Cpu().MilliValue(), tags))
		}
	}

	return ret
}
