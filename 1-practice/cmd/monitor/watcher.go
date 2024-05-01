package main

import (
	"context"
	"fmt"
	"log"
	"monitor/k8s"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type Watcher struct {
	Deployment string
	ClientSet  *kubernetes.Clientset
	MetricsSet *metrics.Clientset
}

func (w *Watcher) Run() {
	var err error
	w.ClientSet, err = k8s.CreateClientSet()
	if err != nil {
		log.Fatal(err)
	}

	w.MetricsSet, err = k8s.CreateMetricsSet()
	if err != nil {
		log.Fatal(err)
	}

	podMetrics, err := w.MetricsSet.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, podMetric := range podMetrics.Items {
		podContainers := podMetric.Containers
		for _, container := range podContainers {
			cpuQuantity, ok := container.Usage.Cpu().AsInt64()
			if !ok {
				log.Fatal("NO!")
				return
			}
			memQuantity, ok := container.Usage.Memory().AsInt64()
			if !ok {
				log.Fatal("NO!!")
				return
			}
			msg := fmt.Sprintf("Container Name: %s \n CPU usage: %d \n Memory usage: %d", container.Name, cpuQuantity, memQuantity)
			fmt.Println(msg)
		}

	}
}
