package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Endpoint struct {
	PodName       string
	PodIP         string
	ContainerName string
	ContainerPort int32
}

var (
	podLatency = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pod_latency",
			Help: "Latency of IPs and ports on pods in seconds",
		},
		[]string{"podname", "containername", "podip", "containerport"},
	)
)

func getEndpoints(kclient *kubernetes.Clientset) []Endpoint {
	var endpoints []Endpoint

	pods, err := kclient.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		if pod.Spec.HostNetwork == false {
			for _, container := range pod.Spec.Containers {
				for _, port := range container.Ports {
					e := new(Endpoint)
					e.PodName = pod.Name
					e.PodIP = pod.Status.PodIP
					e.ContainerName = container.Name
					e.ContainerPort = port.ContainerPort
					endpoints = append(endpoints, *e)
				}
			}
		}
	}
	return endpoints
}

func checkEndpoint(es []Endpoint) {
	var elapsed float64
	var wg sync.WaitGroup

	//As pods die & respawn, every single pod to have ever existed would remain registered on podLatency if no action was taken.
	//Resetting is crude, but solves this. The exporter will only show pods that currently exist, as they get added to the list every time we go through here.
	podLatency.Reset()

	for _, ep := range es {
		wg.Add(1)
		go func(e Endpoint) {
			defer wg.Done()
			log.Println("Attempting to connect to endpoint:", e.PodName, e.ContainerName, e.PodIP, e.ContainerPort)
			start := time.Now()
			_, err := net.DialTimeout("tcp", e.PodIP+":"+strconv.Itoa(int(e.ContainerPort)), time.Duration(1)*time.Second)

			if err != nil {
				log.Println("Connection refused, endpoint unavailable or some other error:", e.PodName, e.ContainerName, e.PodIP, e.ContainerPort)
				elapsed = -1
			} else {
				elapsed = time.Since(start).Seconds()
				log.Println("Endpoint responsse time:", elapsed)
			}

			podLatency.With(prometheus.Labels{"podname": e.PodName, "containername": e.ContainerName, "podip": e.PodIP, "containerport": strconv.Itoa(int(e.ContainerPort))}).Set(elapsed)
		}(ep)
	}
	wg.Wait()
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

	clientset := kubeConnect()
	_, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})

	if err != nil {
		log.Println("Could not list pods")
		http.Error(w, "Could not list pods", 500)
	}

	fmt.Fprintln(w, "this_is_fine.jpeg")
	log.Println("this_is_fine.jpeg")

}

func kubeConnect() *kubernetes.Clientset {
	log.Println("Starting up!")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println("Error getting InClusterConfig")
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("Error getting clientset")
		panic(err.Error())
	}
	return clientset
}

func main() {

	clientset := kubeConnect()

	log.Println("Cluster connection succcessful. Starting endpoint checks")

	go func() {
		for {
			es := getEndpoints(clientset)
			checkEndpoint(es)
			time.Sleep(15 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/healthz", healthCheck)
	http.ListenAndServe(":2113", nil) //Using the default serveMux
}
